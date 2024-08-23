<script>
  import { onDestroy } from "svelte";
  import { current_day_idx, plan } from "../stores.js";
  import { get } from "svelte/store";
  import { GetExerciseById } from "$lib/wailsjs/go/main/App";
  import ExerciseCard from "../../lib/components/ExerciseCard.svelte";

  let current_day;
  let current_exercise_idx = 0;
  let exercises = [];

  plan.subscribe((p) => {
    current_day = p.Days[get(current_day_idx)];

    current_day.ExerciseUnits.map((eu) => {
      GetExerciseById(eu.ExerciseId).then((res) => {
        exercises.push(res);
      });
    });
    exercises = [...exercises];
  });

  let elapsed = 0;
  let elapsed_string = "";
  let last_time = window.performance.now();
  let frame;

  (function loop() {
    frame = requestAnimationFrame(loop);
    const now = window.performance.now();
    elapsed += now - last_time;
    last_time = now;
    elapsed_string = new Date(elapsed).toISOString().substring(11, 19);
  })();

  onDestroy(() => {
    cancelAnimationFrame(frame);
  });
</script>

<div role="group">
  <h2>{current_day.Name} {elapsed_string}</h2>
  <button>Save</button>
</div>
{#if exercises.length === 0}
  <p>No Exercises found</p>
{:else}
  <article>
    <header>
      <div class="center">
        <div class="breathe">
          <h4>{exercises[current_exercise_idx].Name}</h4>
        </div>
        <img
          src={`assets/img/${exercises[current_exercise_idx].Images[0]}`}
          alt={exercises[current_exercise_idx].Name}
        />

        <div role="group">
          {#if current_exercise_idx > 0}
            <button on:click={() => (current_exercise_idx -= 1)}
              >Previous</button
            >
          {/if}
          {#if current_exercise_idx < exercises.length - 1}
            <button on:click={() => (current_exercise_idx += 1)}>Next</button>
          {/if}
        </div>
      </div>
    </header>
    <body>
      <!-- <p>{exercises[current_exercise_idx].Instructions.join("\n")}</p> -->
      {#each current_day.ExerciseUnits[current_exercise_idx].Sets as set, i}
        <ExerciseCard {set} set_idx={i + 1} />
      {/each}
    </body>
  </article>
{/if}

<style>
  .center {
    text-align: center;
    display: flex;
    flex-flow: column;
    justify-content: center;
  }

  .breathe {
    margin-top: 10px;
    margin-bottom: 10px;
  }
</style>
