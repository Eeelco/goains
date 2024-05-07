<script>
  import { onDestroy } from "svelte";
  import { current_day_idx, plan } from "../stores.js";
  import { get } from "svelte/store";
  import { GetExerciseById } from "$lib/wailsjs/go/main/App";

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

<h2>{current_day.Name} {elapsed_string}</h2>
{#if exercises.length === 0}
  <p>No Exercises found</p>
{:else}
  <article>
    <header>
      <p>
        <strong>{exercises[current_exercise_idx].Name}</strong>
      </p>
    </header>
    <body>
      <p>{exercises[current_exercise_idx].Instructions.join("\n")}</p>
      {#if current_exercise_idx > 0}
        <button on:click={() => (current_exercise_idx -= 1)}>Previous</button>
      {/if}
      {#if current_exercise_idx < exercises.length - 1}
        <button on:click={() => (current_exercise_idx += 1)}>Next</button>
      {/if}
    </body>
  </article>
{/if}
