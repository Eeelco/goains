<script>
  import { GetExercises, GetImgPrefix } from "../../lib/wailsjs/go/main/App";
  import { onMount } from "svelte";
  export let modalOpen = false;
  let search_string = "";
  let exercise_list = [];
  let img_prefix;
  onMount(() => {
    GetImgPrefix().then((res) => {
      img_prefix = res;
    });
  });
</script>

<dialog open={modalOpen}>
  <article>
    <h1>Modal</h1>
    <button on:click={() => (modalOpen = false)}>Close</button>

    <input
      type="text"
      bind:value={search_string}
      on:keyup={() => {
        GetExercises(search_string).then((res) => {
          exercise_list = res;
        });
      }}
      placeholder="Exercise name"
    />
    <div class="overflow-auto">
      {#each exercise_list as ex}
        <article>
          <div class="grid">
            <div><h3>{ex.Name}</h3></div>
            <div><button class="secondary">Add</button></div>
          </div>
        </article>
      {/each}
    </div>
  </article>
</dialog>
