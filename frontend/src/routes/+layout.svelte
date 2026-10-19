<script lang="ts">
  import "@wailsio/runtime";
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
        <main class="min-w-0 flex-1 overflow-y-auto">
          <Toaster />
          {@render children()}
        </main>
      </TooltipProvider>
    </div>
  </div>
</SidebarProvider>
