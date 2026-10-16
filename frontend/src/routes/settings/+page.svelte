<script lang="ts">
  import PageHint from "#lib/components/mvz2mm/PageHint.svelte";
  import CardContent from "#lib/components/ui/card/card-content.svelte";
  import CardHeader from "#lib/components/ui/card/card-header.svelte";
  import CardTitle from "#lib/components/ui/card/card-title.svelte";
  import Card from "#lib/components/ui/card/card.svelte";
</script>

<div id="settings-container " class="flex flex-col p-6 gap-4">
  <PageHint title="设置" hint="MVZ2目录、网络源"></PageHint>
  <Card>
    <CardHeader>
      <CardTitle>MVZ2设置</CardTitle>
    </CardHeader>
    <CardContent></CardContent>
  </Card>
</div>
