<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="app">
    <header>
      <h1>Sistema de Gestão Contábil</h1>
      <nav>
        <RouterLink to="/">Abertura de Empresa</RouterLink>
        <RouterLink to="/responsabilidades">Responsabilidades</RouterLink>
      </nav>
    </header>

    <main>
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
.app {
  min-height: 100vh;
}

header {
  background-color: #2c3e50;
  color: white;
  padding: 1.5rem;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

header h1 {
  margin: 0 0 1rem 0;
  font-size: 1.8rem;
  text-align: center;
}

nav {
  display: flex;
  justify-content: center;
  gap: 2rem;
}

nav a {
  color: white;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  transition: background-color 0.3s;
}

nav a:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

nav a.router-link-exact-active {
  background-color: #42b983;
}

main {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}
</style>
