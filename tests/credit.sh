# curl -i -d '{"valor": 1000,"tipo":"c","descricao" : "descricao"}' -H "Content-Type: application/json" -X POST http://localhost:8000/clientes/1/transacoes
curl -i -d '{"valor": 1000,"tipo":"c","descricao" : "descricao"}' -H "Content-Type: application/json" -X POST http://localhost:9999/clientes/1/transacoes
