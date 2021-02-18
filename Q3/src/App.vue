<template>
  <div id="app">
    <h1>Data JSON 1. Search Items Meeting Room</h1>
    <input type="text" v-model="search" placeholder="search Meeting Room">
    <div class="b-card">
      <div v-for="q3 of filteredMr" :key="q3.inventory_id" id="card">
        <h3>{{ q3.name }}</h3>
        <p>{{ q3.type }}</p>
        <p>{{ q3.tags }}</p>
        <p>{{ q3.purchased_at }}</p>
        <p>{{ q3.placement.name }}</p>
      </div>
    </div>

    <h1>Data JSON 2 & 3. Search electronic devices & the furniture</h1>
    <div class="b-card">
      <div v-for="q31 of filteredElc" :key="q31.inventory_id" id="card">
        <h3>{{ q31.name }}</h3>
        <p>{{ q31.type }}</p>
        <p>{{ q31.tags }}</p>
        <p>{{ q31.purchased_at }}</p>
        <p>{{ q31.placement.name }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: 'App',
  data() {
    return {
      q3s: [],
      q3s1: [],
      search: '',
    }
  },
  async created() {
    try {
      const res = await axios.get(`http://localhost:3000/q3`);
      
      this.q3s = res.data;
      this.q3s1 = res.data;
    } catch(e) {
      console.error(e);
    }
  },
  computed: {
    filteredMr:function() {
      return this.q3s.filter((q3)=> {
        return q3.placement.name.match(this.search)
      });
    },
    filteredElc:function() {
      return this.q3s1.filter((q3)=> {
         return q3.type.match(this.search)
      });
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.b-card {
  display: flex;
}
#card {
  background: #eee;
  margin: 10px;
  padding: 5px 10px;
  width: 100%;
}
</style>
