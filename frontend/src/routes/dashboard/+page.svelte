<script lang="ts">
  import PageHint from "#lib/components/mvz2mm/PageHint.svelte";
</script>

<div id="dashboard-container" class="flex flex-col">
  <PageHint title="总览" hint="Mod状态与快捷入口" />
  <div id="status-container" class="flex flex-row">
    <Card></Card>
  </div>
</div>
