O Objetivo do Projeto

Construir uma plataforma de orquestração de vendas flash (Flash Sales) capaz de absorver picos de tráfego extremos (Thundering Herd), garantir a consistência absoluta do inventário (sem overselling - vender o que não existe) e rotear pagamentos de forma resiliente.

## 1. O Problema de Negócio

### O negócio quer faturar dezenas de milhões de dólares em minutos. Para isso, o sistema deve:

- Maximizar o Throughput: Vender todo o estádio no menor tempo possível.
- Garantir Justiça (Fairness): O primeiro a chegar deve ser o primeiro a comprar.
- Risco Zero de Overselling: Se o estádio tem 50.000 lugares, não podemos vender 50.001. A multa por overselling é astronômica e causa um desastre de Relações Públicas (PR).
- Evitar Chargebacks por Falha: Cobrar o cartão do usuário e não emitir o ingresso destrói a margem de lucro por conta de disputas com as adquirentes (Visa/Mastercard).