<script lang="ts">
  import PageHint from "#lib/components/mvz2mm/PageHint.svelte";
  import CardContent from "#lib/components/ui/card/card-content.svelte";
  import CardHeader from "#lib/components/ui/card/card-header.svelte";
  import CardTitle from "#lib/components/ui/card/card-title.svelte";
  import Card from "#lib/components/ui/card/card.svelte";
  import ItemActions from "#lib/components/ui/item/item-actions.svelte";
  import ItemContent from "#lib/components/ui/item/item-content.svelte";
  import ItemHeader from "#lib/components/ui/item/item-header.svelte";
  import Item from "#lib/components/ui/item/item.svelte";
  import { onMount } from "svelte";

  let mvz2Path = $state("None");

  onMount(() => {});
</script>

<div id="settings-container " class="flex flex-col p-6 gap-4">
  <PageHint title="设置" hint="MVZ2目录、网络源"></PageHint>
  <Card>
    <CardHeader class="select-none">
      <CardTitle>MVZ2设置</CardTitle>
    </CardHeader>
    <CardContent>
      <Item variant="outline">
        <ItemContent class="select-none"
          ><ItemHeader>MVZ2路径</ItemHeader>
        </ItemContent>
        <ItemActions>{mvz2Path}</ItemActions>
      </Item>
    </CardContent>
  </Card>
</div>
