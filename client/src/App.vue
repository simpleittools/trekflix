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

</style>
