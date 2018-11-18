package main


import (
	"os"
	"fmt"
	"sort"
	"bufio"
	"strings"
)

type  EmailCount struct {
	email  string
	count  int
}

func main() {
 
	authors,err := os.Open("AUTHORS")
	if err != nil{
		fmt.Println("open AUTHORS file error")
		os.Exit(1)
	}
    
    order,err := os.OpenFile("AUTHORS_BY_SORT", os.O_CREATE | os.O_WRONLY, 0666)
    if err != nil{
    	fmt.Println("open AUTHORS_BY_SORT")
        os.Exit(1)
    }
    defer authors.Close()
    defer order.Close()

    author_map := make(map[string]int)
    var  email_count []EmailCount
    var  email_count_by_email []EmailCount
    var  email_count_by_count []EmailCount
    var temp_email_count EmailCount
    //write_buf := bufio.NewWriter(order)
   
    filesanner := bufio.NewScanner(authors)

    //spit by "@" and orgnized the data as a map
    for filesanner.Scan() {
        line_author := filesanner.Text()
        if strings.Count(line_author, "@") == 1 {
        	temp_email := strings.Replace(strings.Split(line_author, "@")[1],">","",1)
        	author_map[temp_email]++
        }else if strings.Count(line_author, "@") == 2{
            temp_email := strings.Split(line_author, "> <")
            author_map[strings.Split(temp_email[0],"@")[1]]++
            author_map[strings.Replace((strings.Split(temp_email[1], "@"))[1], ">", "",1)]++
        }
    }
    
    //change map to slice for sorted
    for key,count := range author_map{
    	temp_email_count.email = key
    	temp_email_count.count = count
    	email_count = append(email_count, temp_email_count)
    }

    sort.Slice(email_count, func(i,j int) bool {
         return email_count[i].count > email_count[j].count
    })

    for index,element := range email_count {
    	if element.count == 1 {
            email_count_by_count = email_count[:index]
            email_count_by_email = email_count[index:]
            break
    	}
    }

    sort.Slice(email_count_by_email, func(i,j int)bool {
    	return email_count_by_email[i].email < email_count_by_email[j].email
    	
    })
    
    for _,element := range email_count_by_email{
    	temp_email_count.email = element.email
    	temp_email_count.count = element.count
    	email_count_by_count = append(email_count_by_count, temp_email_count)
    }
    /*
    for _,element := range email_count_by_count {
    	//fmt.Println(element.email + " " + string(element.count))
    	_,err := write_buf.WriteString(element.email + " " + string(element.count) + "\n")
    	if err != nil {
    			fmt.Println(err)
    	}
    	write_buf.Flush()
    }
    */
    for index,element := range email_count_by_count {
        format_string := fmt.Sprintf("%d : %-25s %d\n",index, element.email, element.count)
        order.WriteString(format_string)
    }
}