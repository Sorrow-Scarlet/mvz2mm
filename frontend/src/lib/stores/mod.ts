import { writable } from "svelte/store";
import type { ModInfo } from "../../../bindings/mvz2mm/internal/modmanage/models";
import {
  DisableMod,
  EnableMod,
  ListMods,
  RescanPlugins,
  UninstallMod,
} from "../../../bindings/mvz2mm/internal/modmanage/modservice";

export const mods = writable<ModInfo[]>([]);
export const unmanaged = writable<ModInfo[]>([]); // 未托管，只读展示

export async function refreshModList() {
  mods.set(await ListMods());
  unmanaged.set((await RescanPlugins()) ?? []);
}

export async function toggleMod(info: ModInfo) {
  // Enabled以磁盘为准，先调后端成功后再刷新，不乐观更新
  if (info.enabled) {
    await DisableMod(info.id);
  } else {
    await EnableMod(info.id);
  }
  await refreshModList();
}

export async function removeMod(id: string) {
  await UninstallMod(id);
  await refreshModList();
}
