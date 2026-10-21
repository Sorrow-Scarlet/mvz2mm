import { ModInfo } from "../../../bindings/mvz2mm/internal/modmanage/models";
import {
  DisableMod,
  EnableMod,
  GetModStats,
  ListMods,
  RescanPlugins,
  UninstallMod,
} from "../../../bindings/mvz2mm/internal/modmanage/modservice";

class ModsStore {
  mods = $state<ModInfo[]>([]);
  unmanaged = $state<ModInfo[]>([]);
  stats = $state({ total: 0, enabled: 0 });
  async refresh() {
    this.mods = (await ListMods()) ?? [];
    this.unmanaged = (await RescanPlugins()) ?? [];
    this.stats = await GetModStats();
  }

  async toggle(info: ModInfo) {
    if (info.enabled) {
      await DisableMod(info.id);
    } else {
      await EnableMod(info.id);
    }
    await this.refresh();
  }

  async remove(id: string) {
    await UninstallMod(id);
    await this.refresh();
  }
}

export const modsStore = new ModsStore();
