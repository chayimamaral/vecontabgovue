# Sistema de Gestão Contábil - VeContabGo

Aplicativo para gestão contábil desenvolvido em Vue.js que permite:

## Funcionalidades

### 1. Abertura de Empresa
- Cadastro de empresas com informações completas
- Seleção de estado e cidade
- Formatação automática de CNPJ
- Tipos de empresa disponíveis:
  - MEI - Microempreendedor Individual
  - ME - Microempresa
  - EPP - Empresa de Pequeno Porte
  - LTDA - Sociedade Limitada
  - S/A - Sociedade Anônima
  - EIRELI - Empresa Individual de Responsabilidade Limitada

### 2. Geração de Responsabilidades
O sistema gera automaticamente as responsabilidades de acordo com o tipo de empresa:

#### Responsabilidades Fiscais
- Obrigações mensais e anuais
- Declarações específicas por tipo de empresa
- SPED, DCTF, ECF, entre outros

#### Responsabilidades Tributárias
- Impostos federais, estaduais e municipais
- Alíquotas aplicáveis
- IRPJ, CSLL, PIS, COFINS, ICMS, ISS

#### Responsabilidades Contábeis
- Livros contábeis obrigatórios
- Demonstrações financeiras
- Balanço Patrimonial, DRE, SPED Contábil

## Tecnologias Utilizadas

- **Vue.js 3** - Framework JavaScript progressivo
- **TypeScript** - Superset tipado de JavaScript
- **Vite** - Build tool e dev server
- **Vue Router** - Roteamento
- **Pinia** - Gerenciamento de estado

## Como Executar

### Instalação das dependências
```bash
npm install
```

### Executar em modo de desenvolvimento
```bash
npm run dev
```

### Build para produção
```bash
npm run build
```

### Visualizar build de produção
```bash
npm run preview
```

### Executar testes
```bash
npm run test:unit
```

### Lint do código
```bash
npm run lint
```

## Estrutura do Projeto

```
src/
├── assets/          # Arquivos estáticos
├── components/      # Componentes Vue reutilizáveis
├── router/          # Configuração de rotas
├── stores/          # Stores do Pinia
├── views/           # Páginas/Views da aplicação
│   ├── HomeView.vue       # Abertura de Empresa
│   └── AboutView.vue      # Responsabilidades
├── App.vue          # Componente raiz
└── main.ts          # Ponto de entrada da aplicação
```

## Licença

Este projeto é desenvolvido para fins educacionais e de gestão contábil.
