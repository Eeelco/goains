<script>
  import {
    GetExercises,
    GetImgPrefix,
    GetConfig,
  } from "../../lib/wailsjs/go/main/App";
  import { onMount } from "svelte";
  export let modalOpen = false;
  let search_string = "";
  let exercise_list = [];
  export let exercise_lists;
  export let current_day;
  let img_prefix;
  let config;
  onMount(() => {
    GetImgPrefix().then((res) => {
      img_prefix = res;
    });
    GetConfig().then((res) => {
      config = res;
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
            <div>
              <button
                class="secondary"
                on:click={() => {
                  exercise_lists[current_day].push({
                    ExerciseId: ex.Id,
                    ExerciseName: ex.Name,
                    Rest: config.DefaultRest,
                    Sets: [...new Array(config.DefaultNrSets)].map(() => ({
                      Repetitions: config.DefaultNrReps,
                      Weight: 0,
                    })),
                  });
                  exercise_lists = exercise_lists;
                }}>Add</button
              >
            </div>
          </div>
        </article>
      {/each}
    </div>
  </article>
</dialog>
