<script>
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { EventsOn } from "../lib/wailsjs/runtime";
  import { HasPlan } from "../lib/wailsjs/go/main/App";

  let has_plan;
  onMount(() => {
    HasPlan().then((res) => {
      has_plan = res;
    });
  });
</script>

<h1>Main</h1>

<button on:click={() => goto("/create_plan")}>Create new workout plan</button>

{#if has_plan === true}
  <button on:click={() => goto("/workout")}>Continue workout</button>
{:else}
  No current workout
{/if}
