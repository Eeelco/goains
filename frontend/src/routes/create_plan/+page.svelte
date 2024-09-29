<script>
  import { goto } from "$app/navigation";
  import ExercisesList from "./ExercisesList.svelte";
  import { SavePlan } from "$lib/wailsjs/go/backend/App";
  let plan_name = "";
  let plan_description = "";
  let filter_name = "";
  let nr_days = 3;
  let exercise_lists = [...new Array(nr_days)].map(() => []);
  let day_names = Array.from({ length: nr_days }, (_, i) => `Day ${i + 1}`);
  let modalOpen;
  let current_day = 0;

  function renameDay(i) {
    day_names[i] = prompt("New name", day_names[i]) || day_names[i];
  }

  function removeDay(i) {
    exercise_lists = exercise_lists.filter((_, idx) => idx !== i);
    day_names = day_names.filter((_, idx) => idx !== i);
    nr_days -= 1;
  }

  function addExercise(i) {
    current_day = i;
    modalOpen = true;
  }

  function removeExercise(i, exercise) {
    exercise_lists[i] = exercise_lists[i].filter((e) => e !== exercise);
  }

  function savePlan() {
    const plan = {
      Name: plan_name ? plan_name : "New plan",
      Description: plan_description,
      Days: exercise_lists.map((exercises, i) => ({
        Name: day_names[i],
        ExerciseUnits: exercises,
      })),
    };
    SavePlan(plan).then((success) => {
      if (success === true) {
        alert("Plan saved successfully");
        goto("/");
      } else {
        alert(
          "Plan with the same name already exists. Please choose a different name."
        );
      }
    });
  }

  document.addEventListener("keydown", (event) => {
    if (event.key === "Escape" && modalOpen) {
      modalOpen = false;
    }
  });
</script>

<ExercisesList bind:modalOpen bind:exercise_lists bind:current_day />

<div role="group">
  <h1>New Plan</h1>
  <button on:click={savePlan}> Save plan </button>
  <button
    class="secondary"
    on:click={() => {
      goto("/");
    }}
    >Cancel
  </button>
</div>
<input type="text" bind:value={plan_name} placeholder="Plan name" />
<input
  type="text"
  bind:value={plan_description}
  placeholder="Plan description"
/>
<button
  on:click={() => {
    nr_days += 1;
    exercise_lists[nr_days - 1] = [];
    day_names[nr_days - 1] = `Day ${nr_days}`;
  }}>Add day</button
>

<hr />
<div class="grid">
  {#each Array(nr_days) as _, i}
    <div>
      <details role="button" class="outline secondary" open={i == 0 || null}>
        <summary
          >{day_names[i]}<br />{exercise_lists[i].length} Exercises</summary
        >
        <div role="group">
          <button on:click={() => renameDay(i)}>Rename</button>
          <button class="secondary" on:click={() => removeDay(i)}>Delete</button
          >
        </div>
        <button on:click={() => addExercise(i)}>Add exercise</button>
        <div>
          <hr />
          <h3>Exercises</h3>
          {#each exercise_lists[i] as exercise, ex_idx}
            <details>
              <summary
                >{exercise.ExerciseName}<br />{exercise.Sets.length} Sets</summary
              >
              <label>Rest</label>
              <input type="number" bind:value={exercise.Rest} min="1" />
              {#each exercise.Sets as _, j}
                <div role="group">
                  <label
                    >Reps:
                    <input
                      type="number"
                      bind:value={exercise.Sets[j].Repetitions}
                      min="1"
                    /></label
                  >
                  <label
                    >Weight:
                    <input
                      type="number"
                      bind:value={exercise.Sets[j].Weight}
                      min="0"
                      step="0.1"
                    /></label
                  >
                  <a
                    on:click={() => {
                      exercise.Sets = exercise.Sets.filter(
                        (_, idx) => idx !== j
                      );
                      exercise_lists[i] = [...exercise_lists[i]];
                      if (exercise.Sets.length === 0) {
                        removeExercise(i, exercise);
                      }
                    }}>Remove</a
                  >
                </div>
              {/each}
              <button
                on:click={() => {
                  exercise_lists[i][ex_idx].Sets.push({
                    Repetitions: 0,
                    Weight: 0,
                  });
                  exercise_lists[i] = [...exercise_lists[i]];
                }}>Add set</button
              >
              <button
                class="secondary"
                on:click={() => removeExercise(i, exercise)}
                >Remove Exercise</button
              >
            </details>
          {/each}
        </div>
      </details>
    </div>
  {/each}
</div>
