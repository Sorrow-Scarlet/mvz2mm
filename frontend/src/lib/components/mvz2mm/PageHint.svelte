<script lang="ts">
  let { title, hint }: { title: string; hint?: string } = $props();
</script>

<header class="flex flex-col gap-2">
  <h1 class="text-2xl font-semibold tracking-tight">{title}</h1>
  <p>{hint}</p>
</header>
