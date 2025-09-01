# Towel Day API

## Overview
The Towel Day API is a simple API that provides information about Towel Day, celebrated on May 25th in honor of Douglas Adams, the author of "The Hitchhiker's Guide to the Galaxy." This API allows users to check if today is Towel Day and how many days remain until the next Towel Day.

## Running the API in a Podman Container

To run the Towel Day API in a Podman container, follow these steps:

**Install Podman:**
Ensure that Podman is installed on your system. You can refer to the official [Podman documentation](https://podman.io/getting-started/installation) for installation instructions.

**Build the Container:**
Run the following command in the directory containing your `Dockerfile`:

```bash
podman build -t towel-day-api .
```

**Run the Container:**
Start the container with the following command:

```bash
podman run --rm -dt -p 8080:8080 towel-day-api
```

**Access the API:**
You can now access the API by navigating to `http://localhost:8080/is-towel-day` in your web browser or using `curl`.

## API Endpoint

| Method | Endpoint               | Description                          |
|--------|------------------------|--------------------------------------|
| GET    | `/is-towel-day`        | Returns information about Towel Day, including whether today is Towel Day and how many days are left until the next one. |

## Example Query

To check if today is Towel Day, you can use the following command:

```bash
curl -X GET http://localhost:8080/is-towel-day
```

## Example Response

The API will return a JSON response similar to the following:

```json
{
  "is_towel_day": false,
  "towel_day": "May 25",
  "current_date": "September 1",
  "days_until": 237,
  "message": "There are 237 days until Towel Day."
}
```

---

If you have any further questions or need assistance, feel free to ask!

