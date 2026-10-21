<script lang="ts">
  import * as Runtime from "@wailsio/runtime";
  // 强制 esbuild 保留完整 runtime（含 window.js 的拖放注册）
  // 参考 https://github.com/wailsapp/wails/issues/4956
  (window as any).wails = Runtime;
  import "./layout.css";
  import type { LayoutProps } from "./$types";
  import Titlebar from "#lib/shell/Titlebar.svelte";
  import Sidebar from "#lib/shell/Sidebar.svelte";
  import SidebarProvider from "#lib/components/ui/sidebar/sidebar-provider.svelte";
  import { Toaster } from "svelte-sonner";
  import { isDark } from "#lib/stores/dark.svelte.ts";
  import TooltipProvider from "#lib/components/ui/tooltip/tooltip-provider.svelte";
  let { children }: LayoutProps = $props();

  $effect(() => {
    document.documentElement.classList.toggle("dark", isDark());
  });
</script>

<SidebarProvider>
  <div
    id="base-window"
    class="flex flex-col h-svh w-full overflow-hidden bg-background text-foreground"
  >
    <Titlebar />
    <div id="window-content" class="flex min-h-0 flex-1 overflow-hidden">
      <TooltipProvider>
        <Sidebar />
        <main
          data-file-drop-target
          class="min-w-0 flex-1 overflow-y-auto [&.file-drop-target-active]:ring-2 [&.file-drop-target-active]:ring-primary [&.file-drop-target-active]:bg-primary/5
    data-[installing=true]:pointer-events-none data-[installing=true]:opacity-60"
        >
          <Toaster />
          {@render children()}
        </main>
      </TooltipProvider>
    </div>
  </div>
</SidebarProvider>
