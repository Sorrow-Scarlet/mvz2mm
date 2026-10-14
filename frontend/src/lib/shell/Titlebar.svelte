<script lang="ts">
  import { Window } from "@wailsio/runtime";

  let { title = "MVZ2 Mod 管理器" } = $props();
  let maximised = $state(false);

  const toggleMax = async () => {
    if (await Window.IsMaximised()) {
      await Window.Restore();
      maximised = false;
    } else {
      await Window.Maximise();
      maximised = true;
    }
  };
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
  id="titlebar-container"
  class="flex h-9 select-none items-center justify-between border-b border-border bg-background text-foreground [--wails-draggable:drag] shrink-0"
  ondblclick={toggleMax}
>
  <div id="title-container" class="gap-2 pl-3 flex items-center">
    <span id="title-text" class="text-[13px] font-medium">{title}</span>
  </div>

  <div
    id="window-control-container"
    class="flex h-full [--wails-draggable:no-drag]"
  >
    <button
      title="minimize"
      class="w-12 h-full text-muted-foreground transition-colors hover:bg-accent hover:text-accent-foreground"
      onclick={Window.Minimise}
      ><svg class="mx-auto h-3 w-3" viewBox="0 0 12 12">
        <rect x="1" y="5.5" width="10" height="1" fill="currentColor" />
      </svg>
    </button>

    <button
      title="maximise"
      class="w-12 h-full text-muted-foreground transition-colors
             hover:bg-accent hover:text-accent-foreground"
      onclick={toggleMax}
    >
      {#if maximised}
        <svg
          class="mx-auto h-3 w-3"
          viewBox="0 0 12 12"
          fill="none"
          stroke="currentColor"
          stroke-width="1"
        >
          <rect x="3" y="1" width="8" height="8" />
          <rect x="1" y="3" width="8" height="8" class="fill-background" />
        </svg>
      {:else}
        <svg
          class="mx-auto h-3 w-3"
          viewBox="0 0 12 12"
          fill="none"
          stroke="currentColor"
          stroke-width="1"
        >
          <rect x="1.5" y="1.5" width="9" height="9" />
        </svg>
      {/if}
    </button>

    <button
      title="close"
      class="w-12 h-full text-muted-foreground transition-colors
             hover:bg-destructive hover:text-destructive-foreground"
      onclick={Window.Close}
    >
      <svg
        class="mx-auto h-3 w-3"
        viewBox="0 0 12 12"
        fill="none"
        stroke="currentColor"
        stroke-width="1"
      >
        <path d="M1 1 L11 11 M11 1 L1 11" />
      </svg>
    </button>
  </div>
</div>
