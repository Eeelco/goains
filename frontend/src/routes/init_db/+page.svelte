<script>
  import { goto } from "$app/navigation";
  import { EventsOn } from "../../lib/wailsjs/runtime";

  let i = 0;
  let max = 1;
  $: progress = ((i / max) * 100).toFixed(1);

  EventsOn("download_db_progress", (progress) => {
    i = progress[0];
    max = progress[1];
  });
  EventsOn("download_db_done", () => {
    goto("/");
  });
</script>

<h2>Downloading files...</h2>
<progress id="files" value={i} {max}></progress>{progress}%
