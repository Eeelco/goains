<script>
  import { onDestroy } from "svelte";
  import { current_day_idx, plan } from "../stores.js";
  import { get } from "svelte/store";

  let day_idx;
  let exercises;
  let current_day;
  let plan_value;

  plan.subscribe((p) => {
    day_idx = get(current_day_idx);
    console.log("Day idx: ", day_idx);
    plan_value = p;
    console.log("Plan value: ", plan_value);
    current_day = plan_value.Days[day_idx];
    exercises = current_day.Exercises;
  });

  let elapsed = 0;
  let last_time = window.performance.now();
  let frame;

  (function loop() {
    frame = requestAnimationFrame(loop);
    const now = window.performance.now();
    elapsed += now - last_time;
    last_time = now;
  })();

  onDestroy(() => {
    cancelAnimationFrame(frame);
  });
</script>

<h2>Workout {current_day.Name}</h2>
<div>{(elapsed / 1000).toFixed(0)}s</div>
