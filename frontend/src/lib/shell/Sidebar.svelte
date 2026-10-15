<script lang="ts">
  import SidebarContent from "#lib/components/ui/sidebar/sidebar-content.svelte";
  import SidebarFooter from "#lib/components/ui/sidebar/sidebar-footer.svelte";
  import SidebarMenuButton from "#lib/components/ui/sidebar/sidebar-menu-button.svelte";
  import SidebarMenuItem from "#lib/components/ui/sidebar/sidebar-menu-item.svelte";
  import SidebarMenu from "#lib/components/ui/sidebar/sidebar-menu.svelte";
  import { navItems } from "#lib/stores/nav.svelte.ts";
  import { onMount } from "svelte";
  import { GetCurrentVersion } from "../../../bindings/mvz2mm/internal/dashboard/versionservice";
  import { Sidebar } from "../components/ui/sidebar";
  import { page } from "$app/state";
  import { toast } from "svelte-sonner";

  let version = $state("unknown");

  onMount(() => {
    GetCurrentVersion()
      .then((v) => (version = v))
      .catch(() =>
        toast("错误", {
          description: "获取版本号失败",
          action: { label: "x", onClick: () => {} },
        }),
      );
  });
  const copyToClipboard = () => {};
</script>

<Sidebar class="flex px-2 font-medium text-2xl" collapsible="none">
  <SidebarContent class="select-none">
    <SidebarMenu class="mt-2">
      {#each navItems as item (item.title)}
        <SidebarMenuItem>
          <SidebarMenuButton
            isActive={page.url.pathname === item.url}
            class="h-10"
            onclick={() =>
              console.log(`${item.title} clicked; url:${item.url}`)}
          >
            {#snippet child({ props })}<a
                class="flex w-full items-center gap-2"
                href={item.url}
                {...props}
              >
                <item.icon />
                <span>{item.title}</span>
              </a>
            {/snippet}
          </SidebarMenuButton>
        </SidebarMenuItem>
      {/each}
    </SidebarMenu>
  </SidebarContent>
  <SidebarFooter>
    <SidebarMenu>
      <SidebarMenuItem class="flex flex-row h-10 select-none">
        <SidebarMenuButton class="flex-1 shrink-0">i</SidebarMenuButton>
        <SidebarMenuButton
          isActive={false}
          onclick={copyToClipboard}
          class="flex"
        >
          <span>{version}</span>
        </SidebarMenuButton>
      </SidebarMenuItem>
    </SidebarMenu>
  </SidebarFooter>
</Sidebar>
