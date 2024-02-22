<script>
  import { page } from "$app/stores";

  var tabs = [
    { name: "Home", href: "/" },
    { name: "Query", href: "/query" },
  ];

  $: route = "/" + $page.url.pathname.split("/")[1];
</script>

<nav>
  {#each tabs as { name, href }}
    {#if href === route}
      {name}
    {:else}
      <a {href}>{name}</a>
    {/if}
  {/each}
</nav>

<slot />
