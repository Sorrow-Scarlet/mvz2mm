<script>
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

<div
  id="titlebar-container"
  class="flex h-10 select-none items-center justify-between bg-neutral-900 text-neutral-200 [--wails-draggable:drag]"
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
      class="w-12 h-full text-neutral-300 transition-colors hover:bg-neutral-700"
      onclick={Window.Minimise}
      ><svg class="mx-auto h-3 w-3" viewBox="0 0 12 12">
        <rect x="1" y="5.5" width="10" height="1" fill="currentColor" />
      </svg>
    </button>

    <button
      title="maximise"
      class="w-12 h-full text-neutral-300 transition-colors
             hover:bg-neutral-700"
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
          <rect x="1" y="3" width="8" height="8" fill="#171717" />
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
      class="w-12 h-full text-neutral-300 transition-colors
             hover:bg-red-600 hover:text-white"
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
