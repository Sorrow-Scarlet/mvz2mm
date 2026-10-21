import { toast } from "svelte-sonner";
import { InstallModDialog } from "../../../bindings/mvz2mm/internal/modmanage/modservice";
import { modsStore } from "./mod.svelte";

export const handleInstall = async (installing: boolean) => {
  installing = true;
  try {
    const res = await InstallModDialog(); // 来自 bindings
    if (!res) return; // 用户取消
    if (res.installed.length > 0) {
      toast.success(`已安装 ${res.installed.length} 个mod`, {
        description: res.installed.map((m) => m.name).join("、"),
      });
    }
    for (const f of res.failed) toast.error(f);
    await modsStore.refresh();
  } catch (err) {
    toast.error(err instanceof Error ? err.message : String(err));
  } finally {
    installing = false;
  }
};
