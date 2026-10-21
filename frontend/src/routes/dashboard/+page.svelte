<script lang="ts">
  import PageHint from "#lib/components/mvz2mm/PageHint.svelte";
  import CardContent from "#lib/components/ui/card/card-content.svelte";
  import CardDescription from "#lib/components/ui/card/card-description.svelte";
  import CardHeader from "#lib/components/ui/card/card-header.svelte";
  import CardTitle from "#lib/components/ui/card/card-title.svelte";
  import Card from "#lib/components/ui/card/card.svelte";
  import { onMount } from "svelte";
  import { toast } from "svelte-sonner";
  import { GetModStats } from "../../../bindings/mvz2mm/internal/modmanage/modservice";

  let stats = $state({ total: 0, enabled: 0 });
  onMount(async () => {
    stats = await GetModStats();
  });

  let mvz2Info = $state("...");
  const mvz2InfoClicked = () => {
    navigator.clipboard.writeText(mvz2Info);
    toast.info("已成功复制MVZ2信息");
  };
</script>

<div id="dashboard-container" class="flex flex-col p-6 gap-4">
  <PageHint title="总览" hint="Mod状态与快捷入口" />
  <div id="status-container" class="flex flex-row gap-4 select-none">
    <Card class="w-full">
      <CardHeader>
        <CardTitle>{stats.total}</CardTitle>
        <CardDescription>总Mod数</CardDescription>
      </CardHeader>
    </Card>
    <Card class="w-full ">
      <CardHeader>
        <CardTitle>{stats.enabled}</CardTitle>
        <CardDescription>已启用</CardDescription>
      </CardHeader>
    </Card>
  </div>
  <Card class="w-full select-none" onclick={mvz2InfoClicked}>
    <CardHeader>
      <CardTitle>MVZ2信息</CardTitle>
    </CardHeader>
    <CardContent>{mvz2Info}</CardContent>
  </Card>
</div>
