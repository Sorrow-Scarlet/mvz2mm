<script lang="ts">
  import SidebarContent from "#lib/components/ui/sidebar/sidebar-content.svelte";
  import SidebarFooter from "#lib/components/ui/sidebar/sidebar-footer.svelte";
  import SidebarMenuButton from "#lib/components/ui/sidebar/sidebar-menu-button.svelte";
  import SidebarMenuItem from "#lib/components/ui/sidebar/sidebar-menu-item.svelte";
  import SidebarMenu from "#lib/components/ui/sidebar/sidebar-menu.svelte";
  import { nav, navItems } from "#lib/stores/nav.svelte.ts";
  import { onMount } from "svelte";
  import { GetCurrentVersion } from "../../../bindings/mvz2mm/internal/dashboard/versionservice";
  import { Sidebar } from "../components/ui/sidebar";

  let version = $state("unknown");

  onMount(() => {
    GetCurrentVersion()
      .then((v) => (version = v))
      .catch(() => {});
  });
  const changeDark = () => {};
</script>

<Sidebar class="flex px-2 font-medium" collapsible="none">
  <SidebarContent class="select-none">
    <SidebarMenu class="mt-2">
      {#each navItems as item (item.id)}
        <SidebarMenuItem>
          <SidebarMenuButton
            isActive={nav.page == item.id}
            onclick={() => (nav.page = item.id)}
            class="h-10"
          >
            <item.icon />
            <span>{item.title}</span>
          </SidebarMenuButton>
        </SidebarMenuItem>
      {/each}
    </SidebarMenu>
  </SidebarContent>
  <SidebarFooter>
    <SidebarMenu>
      <SidebarMenuItem class="flex flex-col">
        <SidebarMenuButton></SidebarMenuButton>
        <SidebarMenuButton isActive={false} onclick={() => {}}>
          <span class="select-none">{version}</span>
        </SidebarMenuButton>
      </SidebarMenuItem>
    </SidebarMenu>
  </SidebarFooter>
</Sidebar>
