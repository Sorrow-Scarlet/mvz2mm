<script lang="ts">
  import PageHint from "#lib/components/mvz2mm/PageHint.svelte";
  import Button from "#lib/components/ui/button/button.svelte";
  import CardContent from "#lib/components/ui/card/card-content.svelte";
  import CardHeader from "#lib/components/ui/card/card-header.svelte";
  import CardTitle from "#lib/components/ui/card/card-title.svelte";
  import Card from "#lib/components/ui/card/card.svelte";
  import Input from "#lib/components/ui/input/input.svelte";
  import ItemActions from "#lib/components/ui/item/item-actions.svelte";
  import ItemContent from "#lib/components/ui/item/item-content.svelte";
  import ItemDescription from "#lib/components/ui/item/item-description.svelte";
  import ItemHeader from "#lib/components/ui/item/item-header.svelte";
  import Item from "#lib/components/ui/item/item.svelte";
  import { Power, Trash } from "@lucide/svelte";
</script>

<div class="flex flex-col p-6 gap-4">
  <PageHint title="Mod 管理" hint="拖放，启用，删除"></PageHint>
  <Input placeholder="搜索Mod名称..." />

  <Card>
    <CardHeader><CardTitle>已启用</CardTitle></CardHeader>
    <CardContent>
      <Item variant="outline">
        <ItemContent class="select-none">
          <ItemHeader>Mod 名称</ItemHeader>
          <ItemDescription>{`Mod 路径`}</ItemDescription>
        </ItemContent>
        <ItemActions>
          <Button variant="outline"><Power /></Button>
          <Button variant="outline"><Trash /></Button>
        </ItemActions>
      </Item>
    </CardContent>
  </Card>
  <Card><CardHeader><CardTitle>已禁用</CardTitle></CardHeader></Card>
</div>
