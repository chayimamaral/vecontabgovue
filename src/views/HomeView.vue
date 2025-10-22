<script setup lang="ts">
import { ref, computed } from 'vue'

interface Empresa {
  id: number
  nomeEmpresa: string
  cnpj: string
  estado: string
  cidade: string
  tipoEmpresa: string
  atividadePrincipal: string
}

// Estados brasileiros
const estados = [
  { sigla: 'AC', nome: 'Acre' },
  { sigla: 'AL', nome: 'Alagoas' },
  { sigla: 'AP', nome: 'Amapá' },
  { sigla: 'AM', nome: 'Amazonas' },
  { sigla: 'BA', nome: 'Bahia' },
  { sigla: 'CE', nome: 'Ceará' },
  { sigla: 'DF', nome: 'Distrito Federal' },
  { sigla: 'ES', nome: 'Espírito Santo' },
  { sigla: 'GO', nome: 'Goiás' },
  { sigla: 'MA', nome: 'Maranhão' },
  { sigla: 'MT', nome: 'Mato Grosso' },
  { sigla: 'MS', nome: 'Mato Grosso do Sul' },
  { sigla: 'MG', nome: 'Minas Gerais' },
  { sigla: 'PA', nome: 'Pará' },
  { sigla: 'PB', nome: 'Paraíba' },
  { sigla: 'PR', nome: 'Paraná' },
  { sigla: 'PE', nome: 'Pernambuco' },
  { sigla: 'PI', nome: 'Piauí' },
  { sigla: 'RJ', nome: 'Rio de Janeiro' },
  { sigla: 'RN', nome: 'Rio Grande do Norte' },
  { sigla: 'RS', nome: 'Rio Grande do Sul' },
  { sigla: 'RO', nome: 'Rondônia' },
  { sigla: 'RR', nome: 'Roraima' },
  { sigla: 'SC', nome: 'Santa Catarina' },
  { sigla: 'SP', nome: 'São Paulo' },
  { sigla: 'SE', nome: 'Sergipe' },
  { sigla: 'TO', nome: 'Tocantins' }
]

// Cidades por estado (exemplo simplificado)
const cidadesPorEstado: Record<string, string[]> = {
  'SP': ['São Paulo', 'Campinas', 'Santos', 'Ribeirão Preto', 'Sorocaba'],
  'RJ': ['Rio de Janeiro', 'Niterói', 'Petrópolis', 'Volta Redonda', 'Campos dos Goytacazes'],
  'MG': ['Belo Horizonte', 'Uberlândia', 'Contagem', 'Juiz de Fora', 'Betim'],
  'BA': ['Salvador', 'Feira de Santana', 'Vitória da Conquista', 'Camaçari', 'Itabuna'],
  'PR': ['Curitiba', 'Londrina', 'Maringá', 'Ponta Grossa', 'Cascavel'],
  'RS': ['Porto Alegre', 'Caxias do Sul', 'Pelotas', 'Canoas', 'Santa Maria'],
  'PE': ['Recife', 'Jaboatão dos Guararapes', 'Olinda', 'Caruaru', 'Petrolina'],
  'CE': ['Fortaleza', 'Caucaia', 'Juazeiro do Norte', 'Maracanaú', 'Sobral']
}

// Tipos de empresa
const tiposEmpresa = [
  { value: 'MEI', label: 'MEI - Microempreendedor Individual' },
  { value: 'ME', label: 'ME - Microempresa' },
  { value: 'EPP', label: 'EPP - Empresa de Pequeno Porte' },
  { value: 'LTDA', label: 'LTDA - Sociedade Limitada' },
  { value: 'SA', label: 'S/A - Sociedade Anônima' },
  { value: 'EIRELI', label: 'EIRELI - Empresa Individual de Responsabilidade Limitada' }
]

// Form data
const formData = ref({
  nomeEmpresa: '',
  cnpj: '',
  estado: '',
  cidade: '',
  tipoEmpresa: '',
  atividadePrincipal: ''
})

const empresaCadastrada = ref(false)
const empresasCadastradas = ref<Empresa[]>([])

// Computed property para cidades filtradas
const cidadesFiltradas = computed(() => {
  if (!formData.value.estado) return []
  return cidadesPorEstado[formData.value.estado] || ['Cidade não cadastrada - use qualquer nome']
})

// Função para cadastrar empresa
const cadastrarEmpresa = () => {
  if (!formData.value.nomeEmpresa || !formData.value.cnpj || !formData.value.estado || 
      !formData.value.cidade || !formData.value.tipoEmpresa) {
    alert('Por favor, preencha todos os campos obrigatórios')
    return
  }

  empresasCadastradas.value.push({ ...formData.value, id: Date.now() })
  empresaCadastrada.value = true
  
  // Resetar formulário após 2 segundos
  setTimeout(() => {
    formData.value = {
      nomeEmpresa: '',
      cnpj: '',
      estado: '',
      cidade: '',
      tipoEmpresa: '',
      atividadePrincipal: ''
    }
    empresaCadastrada.value = false
  }, 2000)
}

const formatCNPJ = (event: Event) => {
  const input = event.target as HTMLInputElement
  let value = input.value.replace(/\D/g, '')
  
  if (value.length <= 14) {
    value = value.replace(/^(\d{2})(\d)/, '$1.$2')
    value = value.replace(/^(\d{2})\.(\d{3})(\d)/, '$1.$2.$3')
    value = value.replace(/\.(\d{3})(\d)/, '.$1/$2')
    value = value.replace(/(\d{4})(\d)/, '$1-$2')
    formData.value.cnpj = value
  }
}
</script>

<template>
  <div class="home-view">
    <h2>Abertura de Empresa</h2>
    
    <div v-if="empresaCadastrada" class="success-message">
      ✓ Empresa cadastrada com sucesso!
    </div>

    <form @submit.prevent="cadastrarEmpresa" class="company-form">
      <div class="form-group">
        <label for="nomeEmpresa">Nome da Empresa *</label>
        <input 
          type="text" 
          id="nomeEmpresa" 
          v-model="formData.nomeEmpresa"
          required
        />
      </div>

      <div class="form-group">
        <label for="cnpj">CNPJ *</label>
        <input 
          type="text" 
          id="cnpj" 
          v-model="formData.cnpj"
          @input="formatCNPJ"
          placeholder="00.000.000/0000-00"
          maxlength="18"
          required
        />
      </div>

      <div class="form-row">
        <div class="form-group">
          <label for="estado">Estado *</label>
          <select id="estado" v-model="formData.estado" required>
            <option value="">Selecione o estado</option>
            <option v-for="estado in estados" :key="estado.sigla" :value="estado.sigla">
              {{ estado.nome }} ({{ estado.sigla }})
            </option>
          </select>
        </div>

        <div class="form-group">
          <label for="cidade">Cidade *</label>
          <select 
            id="cidade" 
            v-model="formData.cidade" 
            :disabled="!formData.estado"
            required
          >
            <option value="">Selecione a cidade</option>
            <option v-for="cidade in cidadesFiltradas" :key="cidade" :value="cidade">
              {{ cidade }}
            </option>
          </select>
        </div>
      </div>

      <div class="form-group">
        <label for="tipoEmpresa">Tipo de Empresa *</label>
        <select id="tipoEmpresa" v-model="formData.tipoEmpresa" required>
          <option value="">Selecione o tipo de empresa</option>
          <option v-for="tipo in tiposEmpresa" :key="tipo.value" :value="tipo.value">
            {{ tipo.label }}
          </option>
        </select>
      </div>

      <div class="form-group">
        <label for="atividadePrincipal">Atividade Principal</label>
        <input 
          type="text" 
          id="atividadePrincipal" 
          v-model="formData.atividadePrincipal"
          placeholder="Ex: Comércio de produtos alimentícios"
        />
      </div>

      <button type="submit" class="btn-primary">Cadastrar Empresa</button>
    </form>

    <div v-if="empresasCadastradas.length > 0" class="companies-list">
      <h3>Empresas Cadastradas</h3>
      <div class="company-card" v-for="empresa in empresasCadastradas" :key="empresa.id">
        <h4>{{ empresa.nomeEmpresa }}</h4>
        <p><strong>CNPJ:</strong> {{ empresa.cnpj }}</p>
        <p><strong>Tipo:</strong> {{ empresa.tipoEmpresa }}</p>
        <p><strong>Localização:</strong> {{ empresa.cidade }} - {{ empresa.estado }}</p>
        <p v-if="empresa.atividadePrincipal"><strong>Atividade:</strong> {{ empresa.atividadePrincipal }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.home-view {
  max-width: 800px;
  margin: 0 auto;
}

h2 {
  color: #2c3e50;
  margin-bottom: 2rem;
}

.success-message {
  background-color: #42b983;
  color: white;
  padding: 1rem;
  border-radius: 4px;
  margin-bottom: 1rem;
  text-align: center;
  animation: fadeIn 0.3s;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(-10px); }
  to { opacity: 1; transform: translateY(0); }
}

.company-form {
  background-color: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  color: #2c3e50;
  font-weight: 500;
}

input, select {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  transition: border-color 0.3s;
}

input:focus, select:focus {
  outline: none;
  border-color: #42b983;
}

select:disabled {
  background-color: #f5f5f5;
  cursor: not-allowed;
}

.btn-primary {
  background-color: #42b983;
  color: white;
  border: none;
  padding: 1rem 2rem;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
  width: 100%;
}

.btn-primary:hover {
  background-color: #369870;
}

.companies-list {
  margin-top: 3rem;
}

.companies-list h3 {
  color: #2c3e50;
  margin-bottom: 1rem;
}

.company-card {
  background-color: white;
  padding: 1.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  margin-bottom: 1rem;
}

.company-card h4 {
  color: #2c3e50;
  margin-bottom: 0.5rem;
}

.company-card p {
  margin: 0.5rem 0;
  color: #555;
}
</style>
