# Quiz maker app backend

[gin](https://gin-gonic.com/) - web framework
[gorm](https://gorm.io/) - handles database
[render](https://render.com/) - hosting service


## Endpoints

### `POST: /api/v1/quizzes`

example json body:

```json
{
	"createdBy": "Kulothungan",
	"title": "Python Programming",
	"questions": [
		{
			"text": "What is str?",
			"correctOptionIdx": 0,
			"options": [
				{"text": "string"},
				{"text": "integer"},
				{"text": "dict"},
				{"text": "set"}
			]
		},
		{
			"text": "What is dict?",
			"correctOptionIdx": 2,
			"options": [
				{"text": "string"},
				{"text": "integer"},
				{"text": "dict"},
				{"text": "set"}
			]
		}
	]
}
```

example response of the request:

```json
{
	"message": "Quiz created successfully",
	"quidId": 89754522,
}
```

### `GET:  /api/v1/quizzes/{quiz_id}`

example response of the request:

```json
{
	"id": 89754522,
	"createdAt": "2023-10-05T21:39:17.397343371+05:30",
	"title": "Python Programming",
	"createdBy": "Kulothungan",
	"questions": [
		{
			"id": 397813821,
			"quizId": 89754522,
			"options": [
				{
					"id": 635011239,
					"questionId": 397813821,
					"text": "dict"
				},
				{
					"id": 2049870362,
					"questionId": 397813821,
					"text": "string"
				},
				{
					"id": 2724491244,
					"questionId": 397813821,
					"text": "set"
				},
				{
					"id": 3559821882,
					"questionId": 397813821,
					"text": "integer"
				}
			],
			"correctOptionId": 0,
			"text": "What is str?"
		},
		{
			"id": 3766009826,
			"quizId": 89754522,
			"options": [
				{
					"id": 202646390,
					"questionId": 3766009826,
					"text": "integer"
				},
				{
					"id": 1067323032,
					"questionId": 3766009826,
					"text": "string"
				},
				{
					"id": 1419749967,
					"questionId": 3766009826,
					"text": "dict"
				},
				{
					"id": 3492017362,
					"questionId": 3766009826,
					"text": "set"
				}
			],
			"correctOptionId": 0,
			"text": "What is dict?"
		}
	]
}
```


### `POST: /api/v1/submissions`

example json body request

```json
{
	"answeredBy": "Rishvan",
	"quizId": 89754522,
	"answers": [
		{
			"questionId": 397813821,
			"optionId": 2049870362
		},
		{
			"questionId": 3766009826,
			"optionId": 1419749967
		}
	]
}
```

response
```json
{
	"message": "Answer submitted successfully",
	"score": 0
}
```


### `GET:  /api/v1/submissions/{quiz_id}`

example request of the request:

```json
{
	"submissions": [
		{
			"id": 1,
			"quizId": 89754522,
			"score": 1,
			"answeredBy": "Rish",
			"answeredAt": "0001-01-01T00:00:00Z"
		},
		{
			"id": 2,
			"quizId": 89754522,
			"score": 1,
			"answeredBy": "Risdfgh",
			"answeredAt": "0001-01-01T00:00:00Z"
		},
		{
			"id": 3,
			"quizId": 89754522,
			"score": 2,
			"answeredBy": "Risdfgh",
			"answeredAt": "0001-01-01T00:00:00Z"
		},
		{
			"id": 4,
			"quizId": 89754522,
			"score": 0,
			"answeredBy": "Risdfgh",
			"answeredAt": "0001-01-01T00:00:00Z"
		}
	]
}
```
