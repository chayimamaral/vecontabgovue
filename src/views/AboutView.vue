<script setup lang="ts">
import { ref, computed } from 'vue'

// Tipos de empresa para selecionar
const tiposEmpresa = [
  { value: 'MEI', label: 'MEI - Microempreendedor Individual' },
  { value: 'ME', label: 'ME - Microempresa' },
  { value: 'EPP', label: 'EPP - Empresa de Pequeno Porte' },
  { value: 'LTDA', label: 'LTDA - Sociedade Limitada' },
  { value: 'SA', label: 'S/A - Sociedade Anônima' },
  { value: 'EIRELI', label: 'EIRELI - Empresa Individual de Responsabilidade Limitada' }
]

const tipoEmpresaSelecionado = ref('')
const estadoSelecionado = ref('')

// Responsabilidades por tipo de empresa
const responsabilidadesPorTipo: Record<string, any> = {
  'MEI': {
    fiscais: [
      { nome: 'DAS-MEI', periodicidade: 'Mensal', descricao: 'Documento de Arrecadação do Simples Nacional' },
      { nome: 'DASN-SIMEI', periodicidade: 'Anual', descricao: 'Declaração Anual do Simples Nacional' }
    ],
    tributarias: [
      { nome: 'Imposto Fixo Mensal', aliquota: 'Valor fixo', descricao: 'INSS + ICMS/ISS' }
    ],
    contabeis: [
      { nome: 'Relatório Mensal de Receitas', periodicidade: 'Mensal', descricao: 'Controle de faturamento' },
      { nome: 'Guarda de documentos', periodicidade: 'Contínua', descricao: 'Notas fiscais e recibos por 5 anos' }
    ]
  },
  'ME': {
    fiscais: [
      { nome: 'PGDAS-D', periodicidade: 'Mensal', descricao: 'Programa Gerador do Documento de Arrecadação do Simples Nacional' },
      { nome: 'DEFIS', periodicidade: 'Anual', descricao: 'Declaração de Informações Socioeconômicas e Fiscais' },
      { nome: 'EFD-ICMS/IPI', periodicidade: 'Mensal', descricao: 'Escrituração Fiscal Digital' }
    ],
    tributarias: [
      { nome: 'Simples Nacional', aliquota: 'Variável', descricao: 'Regime tributário unificado' },
      { nome: 'ISS', aliquota: '2% a 5%', descricao: 'Imposto Sobre Serviços (se aplicável)' }
    ],
    contabeis: [
      { nome: 'Livro Caixa', periodicidade: 'Diária', descricao: 'Registro de entrada e saída de recursos' },
      { nome: 'Livro Razão', periodicidade: 'Mensal', descricao: 'Registro contábil de operações' },
      { nome: 'Balanço Patrimonial', periodicidade: 'Anual', descricao: 'Demonstração financeira' }
    ]
  },
  'EPP': {
    fiscais: [
      { nome: 'PGDAS-D', periodicidade: 'Mensal', descricao: 'Programa Gerador do Documento de Arrecadação do Simples Nacional' },
      { nome: 'DEFIS', periodicidade: 'Anual', descricao: 'Declaração de Informações Socioeconômicas e Fiscais' },
      { nome: 'EFD-ICMS/IPI', periodicidade: 'Mensal', descricao: 'Escrituração Fiscal Digital' },
      { nome: 'DCTF', periodicidade: 'Mensal', descricao: 'Declaração de Débitos e Créditos Tributários Federais' }
    ],
    tributarias: [
      { nome: 'Simples Nacional', aliquota: 'Variável', descricao: 'Regime tributário unificado' },
      { nome: 'ICMS', aliquota: 'Variável', descricao: 'Imposto sobre Circulação de Mercadorias e Serviços' },
      { nome: 'ISS', aliquota: '2% a 5%', descricao: 'Imposto Sobre Serviços' }
    ],
    contabeis: [
      { nome: 'Livro Diário', periodicidade: 'Diária', descricao: 'Registro cronológico das operações' },
      { nome: 'Livro Razão', periodicidade: 'Mensal', descricao: 'Registro contábil de operações' },
      { nome: 'Balanço Patrimonial', periodicidade: 'Anual', descricao: 'Demonstração financeira' },
      { nome: 'DRE', periodicidade: 'Anual', descricao: 'Demonstração do Resultado do Exercício' }
    ]
  },
  'LTDA': {
    fiscais: [
      { nome: 'SPED Fiscal', periodicidade: 'Mensal', descricao: 'Sistema Público de Escrituração Digital' },
      { nome: 'DCTF', periodicidade: 'Mensal', descricao: 'Declaração de Débitos e Créditos Tributários Federais' },
      { nome: 'ECF', periodicidade: 'Anual', descricao: 'Escrituração Contábil Fiscal' },
      { nome: 'EFD-Contribuições', periodicidade: 'Mensal', descricao: 'PIS e COFINS' }
    ],
    tributarias: [
      { nome: 'IRPJ', aliquota: '15% + 10%', descricao: 'Imposto de Renda Pessoa Jurídica' },
      { nome: 'CSLL', aliquota: '9%', descricao: 'Contribuição Social sobre o Lucro Líquido' },
      { nome: 'PIS', aliquota: '0,65% ou 1,65%', descricao: 'Programa de Integração Social' },
      { nome: 'COFINS', aliquota: '3% ou 7,6%', descricao: 'Contribuição para Financiamento da Seguridade Social' }
    ],
    contabeis: [
      { nome: 'Livro Diário', periodicidade: 'Diária', descricao: 'Registro cronológico das operações' },
      { nome: 'Livro Razão', periodicidade: 'Mensal', descricao: 'Registro contábil de operações' },
      { nome: 'Balanço Patrimonial', periodicidade: 'Trimestral/Anual', descricao: 'Demonstração financeira' },
      { nome: 'DRE', periodicidade: 'Trimestral/Anual', descricao: 'Demonstração do Resultado do Exercício' },
      { nome: 'SPED Contábil', periodicidade: 'Anual', descricao: 'Escrituração Contábil Digital' }
    ]
  },
  'SA': {
    fiscais: [
      { nome: 'SPED Fiscal', periodicidade: 'Mensal', descricao: 'Sistema Público de Escrituração Digital' },
      { nome: 'DCTF', periodicidade: 'Mensal', descricao: 'Declaração de Débitos e Créditos Tributários Federais' },
      { nome: 'ECF', periodicidade: 'Anual', descricao: 'Escrituração Contábil Fiscal' },
      { nome: 'EFD-Contribuições', periodicidade: 'Mensal', descricao: 'PIS e COFINS' },
      { nome: 'DIPJ', periodicidade: 'Anual', descricao: 'Declaração de Informações Econômico-Fiscais' }
    ],
    tributarias: [
      { nome: 'IRPJ', aliquota: '15% + 10%', descricao: 'Imposto de Renda Pessoa Jurídica' },
      { nome: 'CSLL', aliquota: '9%', descricao: 'Contribuição Social sobre o Lucro Líquido' },
      { nome: 'PIS', aliquota: '0,65% ou 1,65%', descricao: 'Programa de Integração Social' },
      { nome: 'COFINS', aliquota: '3% ou 7,6%', descricao: 'Contribuição para Financiamento da Seguridade Social' },
      { nome: 'Dividendos', aliquota: 'Isento', descricao: 'Distribuição aos acionistas' }
    ],
    contabeis: [
      { nome: 'Livro Diário', periodicidade: 'Diária', descricao: 'Registro cronológico das operações' },
      { nome: 'Livro Razão', periodicidade: 'Mensal', descricao: 'Registro contábil de operações' },
      { nome: 'Balanço Patrimonial', periodicidade: 'Trimestral', descricao: 'Demonstração financeira' },
      { nome: 'DRE', periodicidade: 'Trimestral', descricao: 'Demonstração do Resultado do Exercício' },
      { nome: 'SPED Contábil', periodicidade: 'Anual', descricao: 'Escrituração Contábil Digital' },
      { nome: 'Atas de Assembleias', periodicidade: 'Conforme necessário', descricao: 'Registro de decisões societárias' }
    ]
  },
  'EIRELI': {
    fiscais: [
      { nome: 'SPED Fiscal', periodicidade: 'Mensal', descricao: 'Sistema Público de Escrituração Digital' },
      { nome: 'DCTF', periodicidade: 'Mensal', descricao: 'Declaração de Débitos e Créditos Tributários Federais' },
      { nome: 'ECF', periodicidade: 'Anual', descricao: 'Escrituração Contábil Fiscal' },
      { nome: 'EFD-Contribuições', periodicidade: 'Mensal', descricao: 'PIS e COFINS' }
    ],
    tributarias: [
      { nome: 'IRPJ', aliquota: '15% + 10%', descricao: 'Imposto de Renda Pessoa Jurídica' },
      { nome: 'CSLL', aliquota: '9%', descricao: 'Contribuição Social sobre o Lucro Líquido' },
      { nome: 'PIS', aliquota: '0,65% ou 1,65%', descricao: 'Programa de Integração Social' },
      { nome: 'COFINS', aliquota: '3% ou 7,6%', descricao: 'Contribuição para Financiamento da Seguridade Social' }
    ],
    contabeis: [
      { nome: 'Livro Diário', periodicidade: 'Diária', descricao: 'Registro cronológico das operações' },
      { nome: 'Livro Razão', periodicidade: 'Mensal', descricao: 'Registro contábil de operações' },
      { nome: 'Balanço Patrimonial', periodicidade: 'Anual', descricao: 'Demonstração financeira' },
      { nome: 'DRE', periodicidade: 'Anual', descricao: 'Demonstração do Resultado do Exercício' },
      { nome: 'SPED Contábil', periodicidade: 'Anual', descricao: 'Escrituração Contábil Digital' }
    ]
  }
}

const responsabilidades = computed(() => {
  if (!tipoEmpresaSelecionado.value) return null
  return responsabilidadesPorTipo[tipoEmpresaSelecionado.value]
})

const gerarRelatorio = () => {
  if (!tipoEmpresaSelecionado.value) {
    alert('Por favor, selecione um tipo de empresa')
    return
  }
  
  window.print()
}
</script>

<template>
  <div class="about-view">
    <h2>Responsabilidades Fiscais, Tributárias e Contábeis</h2>
    
    <div class="selector-section">
      <div class="form-group">
        <label for="tipoEmpresa">Selecione o Tipo de Empresa</label>
        <select id="tipoEmpresa" v-model="tipoEmpresaSelecionado">
          <option value="">Escolha um tipo de empresa</option>
          <option v-for="tipo in tiposEmpresa" :key="tipo.value" :value="tipo.value">
            {{ tipo.label }}
          </option>
        </select>
      </div>
    </div>

    <div v-if="responsabilidades" class="responsibilities-content">
      <div class="action-bar">
        <button @click="gerarRelatorio" class="btn-secondary">
          📄 Gerar Relatório
        </button>
      </div>

      <div class="responsibility-section">
        <h3>📋 Responsabilidades Fiscais</h3>
        <div class="responsibility-grid">
          <div 
            v-for="(item, index) in responsabilidades.fiscais" 
            :key="index"
            class="responsibility-card"
          >
            <h4>{{ item.nome }}</h4>
            <p class="periodicidade">⏰ {{ item.periodicidade }}</p>
            <p>{{ item.descricao }}</p>
          </div>
        </div>
      </div>

      <div class="responsibility-section">
        <h3>💰 Responsabilidades Tributárias</h3>
        <div class="responsibility-grid">
          <div 
            v-for="(item, index) in responsabilidades.tributarias" 
            :key="index"
            class="responsibility-card"
          >
            <h4>{{ item.nome }}</h4>
            <p class="aliquota">📊 Alíquota: {{ item.aliquota }}</p>
            <p>{{ item.descricao }}</p>
          </div>
        </div>
      </div>

      <div class="responsibility-section">
        <h3>📊 Responsabilidades Contábeis</h3>
        <div class="responsibility-grid">
          <div 
            v-for="(item, index) in responsabilidades.contabeis" 
            :key="index"
            class="responsibility-card"
          >
            <h4>{{ item.nome }}</h4>
            <p class="periodicidade">⏰ {{ item.periodicidade }}</p>
            <p>{{ item.descricao }}</p>
          </div>
        </div>
      </div>
    </div>

    <div v-else class="empty-state">
      <p>👆 Selecione um tipo de empresa acima para visualizar as responsabilidades</p>
    </div>
  </div>
</template>

<style scoped>
.about-view {
  max-width: 1200px;
  margin: 0 auto;
}

h2 {
  color: #2c3e50;
  margin-bottom: 2rem;
}

.selector-section {
  background-color: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  margin-bottom: 2rem;
}

.form-group {
  margin-bottom: 0;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  color: #2c3e50;
  font-weight: 500;
}

select {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
}

select:focus {
  outline: none;
  border-color: #42b983;
}

.action-bar {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 2rem;
}

.btn-secondary {
  background-color: #2c3e50;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s;
}

.btn-secondary:hover {
  background-color: #1a252f;
}

.responsibilities-content {
  animation: fadeIn 0.3s;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.responsibility-section {
  margin-bottom: 3rem;
}

.responsibility-section h3 {
  color: #2c3e50;
  margin-bottom: 1.5rem;
  font-size: 1.5rem;
}

.responsibility-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1rem;
}

.responsibility-card {
  background-color: white;
  padding: 1.5rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  transition: transform 0.3s, box-shadow 0.3s;
}

.responsibility-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0,0,0,0.15);
}

.responsibility-card h4 {
  color: #42b983;
  margin-bottom: 0.5rem;
  font-size: 1.1rem;
}

.periodicidade, .aliquota {
  color: #666;
  font-size: 0.9rem;
  margin: 0.5rem 0;
  font-weight: 500;
}

.responsibility-card p:last-child {
  color: #555;
  margin-top: 0.5rem;
  line-height: 1.5;
}

.empty-state {
  background-color: #f5f5f5;
  padding: 3rem;
  border-radius: 8px;
  text-align: center;
  color: #666;
  font-size: 1.2rem;
}

@media print {
  .selector-section, .action-bar {
    display: none;
  }
  
  .responsibility-card {
    break-inside: avoid;
  }
}
</style>
