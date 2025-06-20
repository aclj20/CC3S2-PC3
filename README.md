
# CC3S2-PC3
## Nombre

Ariana Camila Lopez Julcarima

## Correo

ariana.lopez.j@uni.pe

## Proyecto

Proyecto 6: "DevSecOps local: Escaneo SAST e IaC Security Checks"

## Repositorio grupal:

https://github.com/Akira-13/CC3S2-PC3

## Contribuciones 
Durante el desarrollo del proyecto, implementé la estructura base para gestionar infraestructura como código con Terraform. Automatizé escaneos de seguridad integrando herramientas como Bandit, TFLint y Checkov, además de crear scripts orquestadores, también ersonalicé reglas de validación. Por último, implementé un módulo que genera archivos de configuración que son validados por otro script que implemente de manera automática.


## Instrucciones de uso
```
git clone https://github.com/aclj20/CC3S2-PC3.git

cd iac
terraform init
terraform apply

python scripts/validate_network_json.sh

```