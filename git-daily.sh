#!/bin/bash
# ===========================================================
# Script: git-daily.sh
# Autor: Carlos Amaral
# Descrição: Automatiza o fluxo diário de trabalho com Git
# Repositório: vecontabgovue
# ===========================================================

# Cores para o terminal
GREEN="\033[0;32m"
YELLOW="\033[1;33m"
BLUE="\033[1;34m"
NC="\033[0m" # sem cor

# Verifica se estamos dentro de um repositório Git
if [ ! -d ".git" ]; then
  echo -e "${YELLOW}⚠️  Este diretório não é um repositório Git.${NC}"
  exit 1
fi

# Pergunta uma mensagem de commit
echo -e "${BLUE}📝 Digite a mensagem do commit:${NC}"
read COMMIT_MSG

# Atualiza develop antes de tudo
echo -e "${GREEN}🔄 Atualizando branch develop...${NC}"
git checkout develop >/dev/null 2>&1
git pull origin develop

# Adiciona, commita e envia
echo -e "${GREEN}📦 Adicionando alterações...${NC}"
git add .

if [ -z "$COMMIT_MSG" ]; then
  COMMIT_MSG="Atualizações automáticas"
fi

echo -e "${GREEN}💾 Commitando alterações: '${COMMIT_MSG}'${NC}"
git commit -m "$COMMIT_MSG"

echo -e "${GREEN}☁️ Enviando para o repositório remoto...${NC}"
git push origin develop

# Pergunta se quer promover pra main
echo -e "${BLUE}🚀 Deseja promover para a branch main? (s/n)${NC}"
read RESPOSTA

if [ "$RESPOSTA" == "s" ] || [ "$RESPOSTA" == "S" ]; then
  echo -e "${GREEN}🌿 Mesclando develop → main...${NC}"
  git checkout main
  git pull origin main
  git merge develop
  git push origin main
  git checkout develop
  echo -e "${GREEN}✅ Main atualizada com sucesso!${NC}"
else
  echo -e "${YELLOW}⏸  Alterações mantidas apenas em develop.${NC}"
fi

echo -e "${GREEN}✨ Fluxo concluído com sucesso.${NC}"

