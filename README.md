# Link Shortner API

> A simple link shortner api written in golang using echo web framework

## Routes

- `new` - shorten a new link. Parameter(s): `url`
- `link` - get a redirect using link id. Parameters: `id`

> **Data is stored in a sqlite database for simplicity**

### Core libraries used

- [Echo](https://echo.labstack.com/)
- [Gorm](https://gorm.io/)