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
  import { modsStore } from "#lib/stores/mod.svelte.ts";
  import { Badge, PackageOpen, Power, Trash, Trash2 } from "@lucide/svelte";
  import { onMount } from "svelte";
  import { toast } from "svelte-sonner";
  import type { ModInfo } from "../../../bindings/mvz2mm/internal/modmanage/models";
  import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger,
  } from "#lib/components/ui/alert-dialog/index.ts";
  import CardDescription from "#lib/components/ui/card/card-description.svelte";
  let search = $state("");
  onMount(() => modsStore.refresh());

  const matches = (info: ModInfo) => {
    const q = search.trim().toLowerCase();
    if (!q) return true;
    return (
      info.name.toLowerCase().includes(q) ||
      info.id.toLowerCase().includes(q) ||
      info.author.toLowerCase().includes(q)
    );
  };

  const enabled = $derived(
    modsStore.mods.filter((m) => m.enabled && matches(m)),
  );
  const disabled = $derived(
    modsStore.mods.filter((m) => !m.enabled && matches(m)),
  );
  const handleToggle = async (info: ModInfo) => {
    try {
      await modsStore.toggle(info);
    } catch (err) {
      toast.error(err instanceof Error ? err.message : String(err));
    }
  };
  const handleRemove = async (id: string) => {
    try {
      await modsStore.remove(id);
      toast.success(`成功移除 ${id}`);
    } catch (err) {
      toast.error(err instanceof Error ? err.message : String(err));
    }
  };
</script>

<div class="flex flex-col p-6 gap-4">
  <PageHint title="Mod 管理" hint="拖放，启用，删除"></PageHint>
  <Input placeholder="搜索Mod名称..." />

  {#snippet modItem(info: ModInfo)}
    <Item variant="outline">
      <ItemContent class="select-none">
        <ItemHeader class="flex items-center gap-2">
          {info.name}
          {#if !info.managed}
            <Badge variant="outline">未托管</Badge>
          {/if}
        </ItemHeader>
        <ItemDescription>v{info.version} · {info.author}</ItemDescription>
      </ItemContent>
      <ItemActions>
        <Button
          variant="outline"
          size="icon"
          onclick={() => handleToggle(info)}
          aria-label={info.enabled ? "禁用" : "启用"}
          class={info.enabled ? "text-primary" : "text-muted-foreground"}
        >
          <Power />
        </Button>
        <AlertDialog>
          <AlertDialogTrigger>
            {#snippet child({ props })}
              <Button
                {...props}
                variant="outline"
                size="icon"
                aria-label="卸载"
              >
                <Trash class="text-destructive" />
              </Button>
            {/snippet}
          </AlertDialogTrigger>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>移除 {info.name}？</AlertDialogTitle>
              <AlertDialogDescription>
                将删除其dll与资源文件（含合集共享资源，当该合集最后一个mod被卸载时），此操作不可撤销。
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel>取消</AlertDialogCancel>
              <AlertDialogAction onclick={() => handleRemove(info.id)}
                >移除
              </AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </ItemActions>
    </Item>
  {/snippet}

  <Card>
    <CardHeader>
      <CardTitle
        >已启用 <Badge variant="secondary">{enabled.length}</Badge></CardTitle
      >
    </CardHeader>
    <CardContent class="flex flex-col gap-2">
      {#each enabled as info (info.id)}
        {@render modItem(info)}
      {:else}
        <p class="text-muted-foreground text-sm">
          {search ? "没有匹配的启用mod" : "暂无启用mod，拖拽zip包到窗口安装"}
        </p>
      {/each}
    </CardContent>
  </Card>

  <Card>
    <CardHeader>
      <CardTitle
        >已禁用 <Badge variant="secondary">{disabled.length}</Badge></CardTitle
      >
    </CardHeader>
    <CardContent class="flex flex-col gap-2">
      {#each disabled as info (info.id)}
        {@render modItem(info)}
      {:else}
        <p class="text-muted-foreground text-sm">
          {search ? "没有匹配的禁用mod" : "禁用的mod会显示在这里"}
        </p>
      {/each}
    </CardContent>
  </Card>

  {#if modsStore.unmanaged.length > 0}
    <Card>
      <CardHeader>
        <CardTitle>
          未托管
          <Badge variant="outline">
            <PackageOpen class="size-3" />
            {modsStore.unmanaged.length}
          </Badge>
        </CardTitle>
        <CardDescription
          >plugins目录下存在非本管理器安装的dll，仅展示</CardDescription
        >
      </CardHeader>
      <CardContent class="flex flex-col gap-2">
        {#each modsStore.unmanaged as info (info.id)}
          <Item variant="outline">
            <ItemContent>
              <ItemHeader>{info.name}</ItemHeader>
              <ItemDescription
                >{info.id} · {info.enabled
                  ? "启用中"
                  : "已禁用"}</ItemDescription
              >
            </ItemContent>
          </Item>
        {/each}
      </CardContent>
    </Card>
  {/if}
</div>
