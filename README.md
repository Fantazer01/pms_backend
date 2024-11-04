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

```bash
git clone https://github.com/Fantazer01/pms-backend.git
cd pms-backend
docker build -f ./docker/pms_api/pmsapi.dockerfile -t pms_api .
```
## Launch

```bash
docker run -d -p 8080:8080 pms_api
```

