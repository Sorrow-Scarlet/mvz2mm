<script lang="ts">
  import "@wailsio/runtime";
  import "./layout.css";
  import type { LayoutProps } from "./$types";
  import Titlebar from "#lib/shell/Titlebar.svelte";
  import Sidebar from "#lib/shell/Sidebar.svelte";
  import SidebarProvider from "#lib/components/ui/sidebar/sidebar-provider.svelte";
  import SidebarTrigger from "#lib/components/ui/sidebar/sidebar-trigger.svelte";

  let { children }: LayoutProps = $props();
</script>

<SidebarProvider>
  <div
    id="base-window"
    class="flex flex-col h-svh w-full overflow-hidden bg-background text-foreground"
  >
    <Titlebar />
    <div id="window-content" class="flex min-h-0 flex-1 overflow-hidden">
      <Sidebar />
      <main class="min-w-0 flex-1 overflow-y-auto">
        {@render children()}
      </main>
    </div>
  </div>
</SidebarProvider>
