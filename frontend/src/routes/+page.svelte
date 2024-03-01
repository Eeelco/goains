<script>
  import { goto } from "$app/navigation";
  import { EventsOn } from "../lib/wailsjs/runtime";

  EventsOn("download_db", () => {
    goto("/init_db");
  });
</script>

<h1>Main</h1>
