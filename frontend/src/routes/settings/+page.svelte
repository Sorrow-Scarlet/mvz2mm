<script lang="ts">
  import PageHint from "#lib/components/mvz2mm/PageHint.svelte";
  import Button from "#lib/components/ui/button/button.svelte";
  import CardContent from "#lib/components/ui/card/card-content.svelte";
  import CardHeader from "#lib/components/ui/card/card-header.svelte";
  import CardTitle from "#lib/components/ui/card/card-title.svelte";
  import Card from "#lib/components/ui/card/card.svelte";
  import ItemActions from "#lib/components/ui/item/item-actions.svelte";
  import ItemContent from "#lib/components/ui/item/item-content.svelte";
  import ItemHeader from "#lib/components/ui/item/item-header.svelte";
  import Item from "#lib/components/ui/item/item.svelte";
  import { Folder } from "@lucide/svelte";
  import { onMount } from "svelte";
  import {
    GetGameDir,
    SelectGameDir,
  } from "../../../bindings/mvz2mm/internal/settings/dirservice";
  import { toast } from "svelte-sonner";
  import Tooltip from "#lib/components/ui/tooltip/tooltip.svelte";
  import TooltipTrigger from "#lib/components/ui/tooltip/tooltip-trigger.svelte";
  import TooltipContent from "#lib/components/ui/tooltip/tooltip-content.svelte";

  let mvz2Path = $state("...");

  const refreshPathState = () => {
    GetGameDir()
      .then((d) => {
        mvz2Path = d;
      })
      .then(() => {
        if (mvz2Path == "") {
          mvz2Path = "请手动选择";
          toast.warning(`无法自动查找MVZ2路径，请手动选择`);
        }
      })
      .catch((e) => {
        toast.error(`无法获取MVZ2目录，${e}`);
      });
  };
  const selectGameDir = () => {
    SelectGameDir()
      .then(() => {
        if (mvz2Path == "") {
          toast.info("已取消选择");
        }
        refreshPathState();
      })
      .catch((e) => {
        toast.error(`无法选择指定目录，${e}`);
      });
  };

  onMount(refreshPathState);
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
        <ItemContent>{mvz2Path}</ItemContent>
        <ItemActions>
          <Tooltip>
            <TooltipTrigger>
              {#snippet child({ props })}
                <Button {...props} variant="outline" onclick={selectGameDir}>
                  <Folder />
                </Button>
              {/snippet}
            </TooltipTrigger>
            <TooltipContent>打开文件选择器</TooltipContent>
          </Tooltip>
        </ItemActions>
      </Item>
    </CardContent>
  </Card>
</div>
