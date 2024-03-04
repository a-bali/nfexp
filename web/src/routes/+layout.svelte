<script>
  import { page } from "$app/stores";

  var tabs = [
    { name: "Home", href: "/" },
    { name: "Query", href: "/query" },
    { name: "Reports", href: "/reports" },
  ];

  $: route = "/" + $page.url.pathname.split("/")[1];
</script>

<div class="title">nfexp</div>
<nav>
  <ul class="tabs">
    {#each tabs as { name, href }}
      {#if href === route}
        <li class="active">{name}</li>
      {:else}
        <li class="inactive"><a {href}>{name}</a></li>
      {/if}
    {/each}
  </ul>
</nav>

<slot />

<style>
  :global(html),
  :global(body) {
    position: relative;
    width: 100%;
    height: 100%;
  }

  :global(body) {
    color: #333;
    margin: 0;
    padding: 8px;
    box-sizing: border-box;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
      Oxygen-Sans, Ubuntu, Cantarell, "Helvetica Neue", sans-serif;
  }
  .title {
    float: right;
    font-weight: bold;
    font-size: 150%;
  }
  .tabs {
    margin: 0;
    padding: 0;
    border-bottom: 1px solid #cccccc;
  }
  .tabs li {
    display: inline-block;
    border: 1px solid #cccccc;
    border-bottom: none;
    background-color: #eeeeee;
  }
  li.active {
    background-color: #ffffff;
    position: relative;
    top: 1px;
    font-weight: bold;
    padding: 11px 20px;
  }
  li.inactive {
    border-bottom: 1px solid #eeeeee;
  }
  .tabs li a {
    padding: 10px 20px;
    display: block;
    text-decoration: none;
    color: #000000;
  }
</style>
