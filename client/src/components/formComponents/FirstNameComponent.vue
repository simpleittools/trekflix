<template>
  <div class="mt-1 sm:mt-0 sm:col-span-2">
    <div class="input-frame">
      <span class="input-label"> First Name </span>
      <VeeField type="text" :value="modelValue"
                v-bind="$attrs"
                @input="$emit('update:modelValue', $event.target.value)"
                name="firstName" :rules="(value) => {
                  if(value && value.trim()) return true
                  return 'A First Name is Required'
                }"
                autocomplete="firstName" class="input-field"
      />
    </div>
    <VeeErrorMessage name="firstName" class="form-error"/>
  </div>
</template>

<script>
import {ErrorMessage, Field} from "vee-validate";

export default{
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