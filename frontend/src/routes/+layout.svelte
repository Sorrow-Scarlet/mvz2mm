<script lang="ts">
  import "@wailsio/runtime";
  import "./layout.css";
  import type { LayoutProps } from "./$types";
  import Titlebar from "#lib/shell/Titlebar.svelte";

  let { children }: LayoutProps = $props();
</script>

<div
  id="base-window"
  class="flex h-screen w-screen flex-col bg-neutral-950 text-neutral-100"
>
  <Titlebar />
  <main>{@render children()}</main>
</div>
