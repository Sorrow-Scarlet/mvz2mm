<script>
  import { onMount } from "svelte";
  import { Sidebar } from "../components/ui/sidebar";
  import { GetCurrentVersion } from "../../../bindings/mvz2mm/internal/dashboard/versionservice";
  import { LayoutGrid } from "@lucide/svelte";

  let version = $state("unknown");

  onMount(() => {
    GetCurrentVersion()
      .then((v) => (version = v))
      .catch(() => {});
  });

  const items = [{ id: "overview", title: "总览", icon: LayoutGrid }];
</script>

<Sidebar collapsible="none">{version}</Sidebar>
