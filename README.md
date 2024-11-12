# PMS

PMS - project management system. 
It is course work of student team for 5 course in MEPhI.

The backend of PMS is located here.
Frontend of the project is [here](https://github.com/JustAEro/pms-ui).

Role in project:
- Oleg K. - lead backend developer
- Yuriy Mo. - backend\db developer
- Yuriy Mi. - ML developer
- Iliy P. - lead frontend developer
- Danil S. - technical writer/ frontend developer

## Build

### For product

```bash
git clone https://github.com/Fantazer01/pms-backend.git
cd pms-backend
docker compose -f docker-compose.yml build
```

### For developing/testing

```bash
git clone https://github.com/Fantazer01/pms-backend.git
cd pms-backend
docker compose -f docker-compose.dev.yml build
```

## Launch

Replace **NAME_DOCKER_COMPOSE_FILE** on docker-compose.yml or docker-compose.dev.yml

```bash
docker compose -f {NAME_DOCKER_COMPOSE_FILE} up -d
```

