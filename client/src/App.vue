<template>
  <nav>
    <NavComponent />
  </nav>
  <router-view/>
</template>

<script>
import NavComponent from '@/components/NavComponent'
export default {
  name: 'App',
  components: {
    NavComponent

  }
}
</script>

<style>
body {
  font-family: Arial, Helvetica, sans-serif;
  background-color: #efefee;
}
</style>
