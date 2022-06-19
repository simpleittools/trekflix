<template>

      <span class="input-label"> {{label}} </span>
      <input type="text"
             :value="modelValue"
             @input="$emit('update:modelValue', $event.target.value)"
             v-bind="$attrs"
             class="input-field"/>
</template>

<script setup>
  defineProps({
    label: {
      type: String,
      default: "",
    },
    modelValue: {
      type: String,
      default: "",
    }
  })
</script>

<style scoped>

</style>