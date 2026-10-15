<script lang="ts">
  import PageHint from "#lib/components/mvz2mm/PageHint.svelte";
  import Card from "#lib/components/ui/card/card.svelte";
</script>

<div id="dashboard-container" class="flex flex-col p-6 gap-4">
  <PageHint title="总览" hint="Mod状态与快捷入口" />
  <div id="status-container" class="flex flex-row">
    <Card class="w-full max-w-sm"></Card>
  </div>
</div>
