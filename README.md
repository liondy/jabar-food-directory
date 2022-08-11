# JABAR DIGITAL SERVICE
Create a food service for jabar digital

## How To Run
To run the project, simply use:
```
docker-compose up --build
```

## API Documentation
Returns all available foods in Jabar
```
GET /foods
{
  id: string,
  name: string,
  description: string,
  category: string[]
}
```