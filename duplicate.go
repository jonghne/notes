package main

import (
	"fmt"
	"gopkg.in/resty.v1"
	"time"
)

func main() {
	startAt := time.Now()

	doQuery(q1)
	fmt.Println("q1 use:", time.Since(startAt))
	q2t := time.Now()
	doQuery(q2)
	fmt.Println("q2 use:", time.Since(q2t))
	q3t := time.Now()
	doQuery(q3)
	fmt.Println("q3 use:", time.Since(q3t))
	q4t := time.Now()
	doQuery(q4)
	fmt.Println("q4 use:", time.Since(q4t))
	q5t := time.Now()
	doQuery(q5)
	fmt.Println("q5 use:", time.Since(q5t))
	q6t := time.Now()
	doQuery(q6)
	fmt.Println("q6 use:", time.Since(q6t))
	q7t := time.Now()
	doQuery(q7)
	fmt.Println("q7 use:", time.Since(q7t))
	q8t := time.Now()
	doQuery(q8)
	fmt.Println("q8 use:", time.Since(q8t))

	fmt.Println("finish use:", time.Since(startAt))
}

func doQuery(body string) {
	resp, err := resty.R().SetHeaders(map[string]string{
		"Content-Type":"application/json",
	}).SetBody(body).Post("http://172.16.5.122:9200/approval_dev/duplicate_feature_data/_search")
	if err != nil {
		panic(err)
	}
	if resp.StatusCode() != 200 {
		panic("invalid resp:" + string(resp.Body()))
	}
}


var (
	q1 = `
{
  "query": {
    "bool": {
      "must_not": { "match": { "order_no": "111_155866806554430097888973"  }},
      "should": [
        { "match": { "immediate_contact_phone": "bjjm68cdb5iojhf6g9g0" }},
        { "match": { "other_contact_phone": "bjjm68cdb5iojhf6g9"   }}
      ]
    }
  }
}
`
	q2 = `
{
  "query": {
    "terms": {
      "apply_phone": [
        "1231741841", "bjjm68cdb5iojhf6g9fg"
      ]
    }
  }
}
`
	q3 = `
{
  "query": {
    "bool": {
      "should": [
        { "terms": { "immediate_contact_phone": ["bjjm68cdb5iojhf6g9g0"] }},
        { "terms": { "other_contact_phone": ["bjjm68cdb5iojhf6g9", "1874812471"]   }}
      ]
    }
  }
}
`
	q4 = `
{
  "query": {
    "bool": {
      "should": [
        { 
        	"bool": {
	        	"must": [ 
		        	{"match": {"id_type": "PanCard"}}, 
		        	{"match": {"id_card_num": "1548781231afas"}}
	        	]	
        	}
        },
        { 
        	"bool": {
	        	"must": [ 
		        	{"match": {"addr_card_type": "Addr"}}, 
		        	{"match": {"addr_card_num": "123123213151"}}
	        	]	
        	}
        }
      ]
    }
  }
}
`
	q5 = `
{
  "query": {
    "bool": {
      "must_not": {
        "match": {
          "order_no": "111_155866806554430097888973"
        }
      },
      "should": [
        {
          "bool": {
            "must_not": [ { "match": { "id_type": "PanCard" } } ],
            "must": [ { "match": { "id_card_num": "1548781231afas" } } ]
          }
        },
        {
          "bool": {
            "must_not": [ { "match": { "addr_card_type": "Addr" } } ],
            "must": [ { "match": { "addr_card_num": "123123213151" } } ]
          }
        }
      ]
    }
  }
}
`
	q6 = `
{
  "query": {
    "bool": {
      "must_not": {
        "match": {
          "order_no": "111_155866806554430097888973"
        }
      },
      "must": [
        {
          "match": {
            "full_name": "acbj cnzjk"
          }
        }
      ]
    }
  }
}
`
	q7 = `
{
  "query": {
    "bool": {
      "must_not": {
        "match": {
          "order_no": "111_155866806554430097888973"
        }
      },
      "must": [
        {
          "match": {
            "device_mac": "17cbxzbuyh124"
          }
        }
      ]
    }
  }
}
`
	q8 = `
{
  "query": {
    "bool": {
      "must_not": {
        "match": {
          "order_no": "111_155866806554430097888973"
        }
      },
      "must": [
        {
          "match": {
            "apply_phone": "18y78y72131"
          }
        }
      ]
    }
  }
}
`
)