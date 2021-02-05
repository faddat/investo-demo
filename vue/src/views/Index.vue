<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form path="faddat.investodemo.investodemo" type="company" :fields="[ ['creator', 1, 'string'] , ['mktgroi', 2, 'string'] , ['logo', 3, 'string'] , ['netflow', 4, 'string'] , ['industry', 5, 'string'] ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
