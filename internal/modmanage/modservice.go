package modmanage

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mvz2mm/internal/settings"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	disabledSuffix = ".disabled"
	modManifest    = "manifest.json"
	packManifest   = "mvz2pack.json"
)

type FileType string

const (
	typeDll      FileType = "dll"
	typeResource FileType = "resource"
	typeOther    FileType = "other"
)

type ModInfo struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Author       string   `json:"author"`
	Description  string   `json:"description"`
	Dependencies []string `json:"dependencies"`
	Enabled      bool     `json:"enabled"`
	Managed      bool     `json:"managed"`
}

type RawModManifest struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Author       string   `json:"author"`
	Description  string   `json:"description"`
	Dependencies []string `json:"dependencies"`
}
type RawPackManifest struct {
	Name       string   `json:"Name"`
	Author     string   `json:"Author"`
	CreatedUtc string   `json:"CreatedUtc"`
	Mods       []string `json:"Mods"`
}
type InstalledFile struct {
	Path string   `json:"path"` // 使用`/`，直接拼接
	Type FileType `json:"type"`
}
type ModRecord struct {
	Info        ModInfo         `json:"info"`
	Files       []InstalledFile `json:"files"`
	SourcePack  string          `json:"sourcePack,omitempty"`
	InstalledAt time.Time       `json:"installedAt"`
	IsPackRes   bool            `json:"isPackRes,omitempty"`
}

type ModService struct {
	dir      *settings.DirService
	registry map[string]*ModRecord
}

func NewModService(dir *settings.DirService) *ModService {
	s := &ModService{dir: dir, registry: map[string]*ModRecord{}}
	s.loadRegistry()
	return s
}

// ModStats dashboard用的mod统计数据
type ModStats struct {
	Total   int `json:"total"`   // 托管mod总数
	Enabled int `json:"enabled"` // 当前启用数
}

// GetModStats 返回托管mod的总数与启用数，Enabled以磁盘实际状态为准
func (m *ModService) GetModStats() ModStats {
	var total, enabled int
	for _, rec := range m.registry {
		if rec.IsPackRes {
			continue
		}
		total++
		for _, f := range rec.Files {
			if f.Type == typeDll && m.fileExists(f.Path) {
				enabled++
				break
			}
		}
	}
	return ModStats{Total: total, Enabled: enabled}
}

// PendingFile zip => 目标相对路径 + 归属
type PendingFile struct {
	zf    *zip.File
	rel   string
	ft    FileType
	owner string // modID；空串 = 合集共享资源
}

// region 前端

func (m *ModService) InstallMod(zipPath string) (*ModInfo, error) {
	if !m.dir.HasGameDirSet() {
		return nil, errors.New("请先在设置页选择游戏目录")
	}
	if !strings.EqualFold(filepath.Ext(zipPath), ".zip") {
		return nil, errors.New("仅支持zip格式的mod包")
	}

	zr, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, errors.New("无法读取压缩包，请确认为zip格式")
	}
	defer zr.Close()

	// 靠描述文件名分流：mvz2pack.json=合集，manifest.json=单mod
	kind, manifestPath, err := findManifest(&zr.Reader)
	if err != nil {
		return nil, err
	}
	if kind == packManifest {
		var raw RawPackManifest
		if err := decodeManifest(&zr.Reader, manifestPath, &raw); err != nil {
			return nil, err
		}
		return m.installCollection(&zr.Reader, manifestPath, &raw)
	}
	var raw RawModManifest
	if err := decodeManifest(&zr.Reader, manifestPath, &raw); err != nil {
		return nil, err
	}
	return m.installSingle(&zr.Reader, manifestPath, &raw)
}

// DisableMod 重命名dll为.disabled，资源与伴随文件不动
func (m *ModService) DisableMod(modID string) error {
	rec, ok := m.registry[modID]
	if !ok || rec.IsPackRes {
		return errors.New("mod不存在: " + modID)
	}
	for _, f := range rec.Files {
		if f.Type != typeDll || m.fileExists(f.Path+disabledSuffix) {
			continue
		}
		from := m.absPath(f.Path)
		if err := os.Rename(from, from+disabledSuffix); err != nil {
			return errors.New("禁用失败，请确认游戏已关闭: " + err.Error())
		}
	}
	return nil
}

// EnableMod 将.disabled恢复为.dll
func (m *ModService) EnableMod(modID string) error {
	rec, ok := m.registry[modID]
	if !ok || rec.IsPackRes {
		return errors.New("mod不存在: " + modID)
	}
	for _, f := range rec.Files {
		if f.Type != typeDll {
			continue
		}
		disabled := m.absPath(f.Path) + disabledSuffix
		if _, err := os.Stat(disabled); err != nil {
			continue
		}
		if err := os.Rename(disabled, disabled[:len(disabled)-len(disabledSuffix)]); err != nil {
			return errors.New("启用失败，请确认游戏已关闭: " + err.Error())
		}
	}
	return nil
}

// UninstallMod 按注册表清单逆序删除；合集最后一个mod卸载时连带删除共享资源
func (m *ModService) UninstallMod(modID string) error {
	rec, ok := m.registry[modID]
	if !ok || rec.IsPackRes {
		return errors.New("mod不存在: " + modID)
	}

	for i := len(rec.Files) - 1; i >= 0; i-- {
		full := m.absPath(rec.Files[i].Path)
		if err := os.Remove(full); err != nil && !os.IsNotExist(err) {
			return errors.New("卸载失败，请确认游戏已关闭: " + err.Error())
		}
		m.removeEmptyDirs(filepath.Dir(full))
	}

	pack := rec.SourcePack
	delete(m.registry, modID)

	if pack != "" {
		remaining, resRecID := 0, ""
		for id, r := range m.registry {
			if r.SourcePack != pack {
				continue
			}
			if r.IsPackRes {
				resRecID = id
			} else {
				remaining++
			}
		}
		if remaining == 0 && resRecID != "" {
			for i := len(m.registry[resRecID].Files) - 1; i >= 0; i-- {
				full := m.absPath(m.registry[resRecID].Files[i].Path)
				if err := os.Remove(full); err != nil && !os.IsNotExist(err) {
					slog.Warn("清理合集资源失败", "path", full, "err", err)
				}
			}
			m.removeEmptyDirs(m.absPath(filepath.ToSlash(filepath.Dir(m.registry[resRecID].Files[0].Path))))
			delete(m.registry, resRecID)
		}
	}

	m.saveRegistry()
	return nil
}

// ListMods 返回所有托管mod，Enabled以磁盘实际状态为准，比较耗时
func (m *ModService) ListMods() []ModInfo {
	var out []ModInfo
	for _, rec := range m.registry {
		if rec.IsPackRes {
			continue
		}
		info := rec.Info
		info.Enabled = false
		for _, f := range rec.Files {
			if f.Type == typeDll && m.fileExists(f.Path) {
				info.Enabled = true
			}
		}
		out = append(out, info)
	}
	return out
}

// RescanPlugins 扫描plugins下未被托管的dll, 及其耗时的操作(取决于装了多少mod)
func (m *ModService) RescanPlugins() []ModInfo {
	entries, err := os.ReadDir(m.dir.GetModDllDir())
	if err != nil {
		return nil
	}
	owned := map[string]bool{}
	for _, rec := range m.registry {
		for _, f := range rec.Files {
			owned[f.Path] = true
		}
	}
	var unmanaged []ModInfo
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || (!strings.HasSuffix(name, ".dll") && !strings.HasSuffix(name, disabledSuffix)) {
			continue
		}
		if owned[settings.ExportmodDllDir+"/"+name] {
			continue
		}
		unmanaged = append(unmanaged, ModInfo{
			ID:      name,
			Name:    strings.TrimSuffix(strings.TrimSuffix(name, disabledSuffix), ".dll"),
			Enabled: !strings.HasSuffix(name, disabledSuffix),
			Managed: false,
		})
	}
	return unmanaged
}

// region PRIVATE

// absPath 接收相对路径rel，与GameDir拼接。 rel必须是正斜杠划分的版本
func (m *ModService) absPath(rel string) string {
	return filepath.Join(m.dir.GameDir, filepath.FromSlash(rel))
}

// fileExists通过相对路径检查文件否存在，返回T/F
func (m *ModService) fileExists(rel string) bool {
	_, err := os.Stat(m.absPath(rel))
	return err == nil
}

// ownerOf根据rel文件名查找对应Mod记录。从注册表的每个mod的每个文件查找，速度On(总表长度)套On(单个mod文件表长度)
func (m *ModService) ownerOf(rel string) *ModRecord {
	for _, rec := range m.registry {
		for _, f := range rec.Files {
			if f.Path == rel {
				return rec
			}
		}
	}
	return nil
}

// removeEmptyDirs接收目录，检查为空就移除；On操作
func (m *ModService) removeEmptyDirs(dir string) {
	cleanGame := filepath.Clean(m.dir.GameDir)
	for filepath.Clean(dir) != cleanGame {
		if os.Remove(dir) != nil {
			return // 非空即停
		}
		// 更新dir，用于继续遍历
		dir = filepath.Dir(dir)
	}
}

// 在用户的APPDATA目录或跨平台数据目录保存本mod管理器的数据
func (m *ModService) registryPath() string {
	cfg, _ := os.UserConfigDir()
	return filepath.Join(cfg, "MVZ2ModManager", "mods.json")
}

// 保存并更新保存的数据
func (m *ModService) saveRegistry() {
	p := m.registryPath()
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		slog.Error("save registry", "err", err)
		return
	}
	records := make([]*ModRecord, 0, len(m.registry))
	for _, rec := range m.registry {
		records = append(records, rec)
	}
	data, _ := json.MarshalIndent(records, "", "  ")
	if err := os.WriteFile(p, data, 0o644); err != nil {
		slog.Error("save registry", "err", err)
	}
}

func (m *ModService) loadRegistry() {
	data, err := os.ReadFile(m.registryPath())
	if err != nil {
		slog.Error("load modmanager registry", "err", err)
		return
	}
	var records []*ModRecord
	if json.Unmarshal(data, &records) == nil {
		for _, rec := range records {
			m.registry[rec.Info.ID] = rec
		}
	}
}

// hashString 接收字符串，并计算字符串短哈希(16位)
func hashString(str string) string {
	// data必须是字符数组，需要对str转换
	sum := sha256.Sum256([]byte(str))
	// EncodeToString要求传入切片类型，返回字符串，随后仅截取16位返回
	return hex.EncodeToString(sum[:])[:16]
}

// commit 冲突检查后逐文件落盘（防zip slip越界）
func (m *ModService) commit(pendings []PendingFile) error {
	for _, p := range pendings {
		if m.fileExists(p.rel) && m.ownerOf(p.rel) == nil {
			return errors.New("文件冲突: " + p.rel + " 已存在，请先卸载对应mod或手动清理")
		}
	}
	for _, p := range pendings {
		if err := m.extractFile(p.zf, p.rel); err != nil {
			return fmt.Errorf("解压失败 %s: %w", p.rel, err)
		}
	}
	return nil
}

func (m *ModService) extractFile(zf *zip.File, relTarget string) error {
	target := m.absPath(relTarget)
	cleanGame := filepath.Clean(m.dir.GameDir) + string(filepath.Separator)
	if !strings.HasPrefix(filepath.Clean(target)+string(filepath.Separator), cleanGame) {
		return errors.New("压缩包内含非法路径: " + relTarget)
	}
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return err
	}
	src, err := zf.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(target)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	return err
}

// findManifest 定位压缩包内层级最浅的描述文件（兼容多包一层目录），
// 返回描述文件种类、路径。
func findManifest(zr *zip.Reader) (string, string, error) {
	best := ""
	for _, f := range zr.File {
		base := strings.ToLower(filepath.Base(f.Name))
		if base != modManifest && base != packManifest {
			continue
		}
		if best == "" || strings.Count(f.Name, "/") < strings.Count(best, "/") {
			best = f.Name
		}
	}
	if best == "" {
		return "", "", errors.New("压缩包内未找到 manifest.json 或 mvz2pack.json")
	}
	if strings.EqualFold(filepath.Base(best), packManifest) {
		return packManifest, best, nil
	}
	return modManifest, best, nil
}

func decodeManifest(zr *zip.Reader, path string, v any) error {
	f, err := zr.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(v); err != nil {
		return errors.New("描述文件解析失败: " + err.Error())
	}
	return nil
}

// rootPrefix 剥掉zip包裹目录前缀，如 "SomeFolder/" → ""
func rootPrefix(manifestPath string) string {
	if i := strings.LastIndex(manifestPath, "/"); i >= 0 {
		return manifestPath[:i+1]
	}
	return ""
}

// scanPack 按标准结构扫描zip内文件，分类为dll/伴随文件/资源。
// 两个前缀之外的路径不符合Protocol，警告并跳过。
func scanPack(zr *zip.Reader, prefix string) (dlls, others, resources []PendingFile) {
	for _, f := range zr.File {
		if f.FileInfo().IsDir() {
			continue
		}
		rel := strings.TrimPrefix(filepath.ToSlash(f.Name), prefix)
		switch {
		case rel == modManifest || rel == packManifest:
			// 描述文件不落盘
		case strings.HasPrefix(rel, settings.ExportmodDllDir+"/"):
			pf := PendingFile{zf: f, rel: rel}
			if strings.HasSuffix(rel, ".dll") {
				pf.ft = typeDll
				dlls = append(dlls, pf)
			} else {
				pf.ft = typeOther
				others = append(others, pf)
			}
		case strings.HasPrefix(rel, settings.Exportmvz2DataDir+"/"):
			resources = append(resources, PendingFile{zf: f, rel: rel, ft: typeResource})
		default:
			slog.Warn("包内不符合Protocol的路径，跳过", "path", rel)
		}
	}
	return
}

// installCollection 安装合集包。
// Mods[i] ↔ plugins/<名字>.dll 一一对应实现单独启停；
// 资源无mod级映射，挂合集共享记录，最后一个mod卸载时删除。
func (m *ModService) installCollection(zr *zip.Reader, manifestPath string, raw *RawPackManifest) (*ModInfo, error) {
	prefix := rootPrefix(manifestPath)
	packID := hashString(raw.Name + raw.CreatedUtc)

	dlls, others, resources := scanPack(zr, prefix)

	// 归属分配：Mods[i]认领同名dll与伴随文件，认领不到的跳过
	var pendings []PendingFile
	claimed := map[string]bool{}
	var firstInfo *ModInfo
	for _, modName := range raw.Mods {
		var files []InstalledFile
		for _, d := range dlls {
			base := strings.TrimSuffix(filepath.Base(d.rel), ".dll")
			if strings.EqualFold(base, modName) {
				pendings = append(pendings, PendingFile{d.zf, d.rel, d.ft, modName})
				files = append(files, InstalledFile{d.rel, d.ft})
				claimed[d.rel] = true
			}
		}
		for _, o := range others {
			base := strings.TrimSuffix(filepath.Base(o.rel), filepath.Ext(o.rel))
			if strings.EqualFold(base, modName) {
				pendings = append(pendings, PendingFile{o.zf, o.rel, o.ft, modName})
				files = append(files, InstalledFile{o.rel, o.ft})
				claimed[o.rel] = true
			}
		}
		if len(files) == 0 {
			slog.Warn("合集声明的mod在plugins下没有对应dll", "mod", modName)
			continue
		}
		rec := &ModRecord{
			Info: ModInfo{
				ID:           modName,
				Name:         modName,
				Version:      raw.CreatedUtc, // 合集无版本字段，用时间戳占位
				Author:       raw.Author,
				Dependencies: []string{},
				Managed:      true,
			},
			Files:       files,
			SourcePack:  packID,
			InstalledAt: time.Now(),
		}
		m.registry[modName] = rec
		if firstInfo == nil {
			info := rec.Info
			info.Enabled = true
			firstInfo = &info
		}
	}
	if firstInfo == nil {
		return nil, errors.New("合集中没有任何一个mod能在plugins下找到对应dll")
	}
	// 未被任何Mods[i]认领的dll/伴随文件：不落盘，避免产生注册表外孤儿文件
	for _, p := range append(dlls, others...) {
		if !claimed[p.rel] {
			slog.Warn("合集包内存在Mods未声明的文件，跳过", "path", p.rel)
		}
	}
	for _, r := range resources {
		pendings = append(pendings, PendingFile{r.zf, r.rel, r.ft, ""})
	}

	if err := m.commit(pendings); err != nil {
		return nil, err
	}

	// 共享资源挂隐藏记录
	if len(resources) > 0 {
		resID := "__pack_res_" + packID
		files := make([]InstalledFile, len(resources))
		for i, r := range resources {
			files[i] = InstalledFile{r.rel, r.ft}
		}
		m.registry[resID] = &ModRecord{
			Info:        ModInfo{ID: resID, Name: "合集资源: " + raw.Name},
			Files:       files,
			SourcePack:  packID,
			IsPackRes:   true,
			InstalledAt: time.Now(),
		}
	}

	m.saveRegistry()
	return firstInfo, nil
}

// installSingle 安装单mod包。
// 包内按标准结构落盘，所有dll/伴随文件/资源都归这一个mod。
func (m *ModService) installSingle(zr *zip.Reader, manifestPath string, raw *RawModManifest) (*ModInfo, error) {
	if raw.ID == "" {
		return nil, errors.New("manifest缺少id字段")
	}
	prefix := rootPrefix(manifestPath)

	dlls, others, resources := scanPack(zr, prefix)
	if len(dlls) == 0 {
		return nil, errors.New(settings.ExportmodDllDir + " 下未找到任何dll文件")
	}

	var pendings []PendingFile
	var files []InstalledFile
	claim := func(list []PendingFile) {
		for _, p := range list {
			pendings = append(pendings, PendingFile{p.zf, p.rel, p.ft, raw.ID})
			files = append(files, InstalledFile{p.rel, p.ft})
		}
	}
	claim(dlls)
	claim(others)
	claim(resources)

	if err := m.commit(pendings); err != nil {
		return nil, err
	}

	rec := &ModRecord{
		Info: ModInfo{
			ID:           raw.ID,
			Name:         raw.Name,
			Version:      raw.Version,
			Author:       raw.Author,
			Description:  raw.Description,
			Dependencies: raw.Dependencies,
			Managed:      true,
		},
		Files:       files,
		InstalledAt: time.Now(),
	}
	m.registry[raw.ID] = m.registry[raw.ID] // 占位防误删，实际下一行覆盖
	m.registry[raw.ID] = rec
	m.saveRegistry()

	info := rec.Info
	info.Enabled = true
	return &info, nil
}
