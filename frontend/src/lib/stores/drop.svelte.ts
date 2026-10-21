import { Events } from "@wailsio/runtime";
import { toast } from "svelte-sonner";
import { InstallMod } from "../../../bindings/mvz2mm/internal/modmanage/modservice";
import { refreshModList } from "./mod";

const DropEvent = "files-dropped";
let installing = false;

export function setupModDropZone(mainEl: HTMLElement) {
  Events.On(DropEvent, async (ev: any) => {
    const raw = ev?.data;
    const paths: string[] = Array.isArray(raw?.[0]) ? raw[0] : (raw ?? []);
    if (paths.length === 0) return;

    if (installing) {
      toast.error("正在安装上一个mod，请稍候");
      return;
    }

    installing = true;
    mainEl.dataset.installing = "true";
    try {
      for (const p of paths) {
        try {
          const info = await InstallMod(p);
          toast.success(`安装成功：${info?.name ?? p}`);
          await refreshModList();
        } catch (err) {
          toast.error(
            `安装失败：${err instanceof Error ? err.message : String(err)}`,
          );
        }
      }
    } finally {
      installing = false;
      delete mainEl.dataset.installing;
    }
  });
}
