<template>
  <div class="mt-1 sm:mt-0 sm:col-span-2">
    <div class="input-frame">
      <span class="input-label"> Username </span>
      <VeeField type="text" :value="modelValue"
                v-bind="$attrs"
                @input="$emit('update:modelValue', $event.target.value)"
                name="username" :rules="(value) => {
                  if(value && value.trim()) return true
                  return 'The Username Field is Required'
                }"
                autocomplete="username" class="input-field"
      />
    </div>
    <VeeErrorMessage name="username" class="form-error"/>
  </div>
</template>

<script>
import {ErrorMessage, Field} from "vee-validate";

export default{
  inheritAttrs: false,
  components: {
    VeeField: Field,
    VeeErrorMessage: ErrorMessage,

  },
  }
</script>

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