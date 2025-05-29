package day4

/*

const filename = "my_test_file.txt"
	contents := []string{"Hello Go!", " this is my first file."}

	for _, content := range contents {
		day4.AppendToFile(filename, content)
	}

	data, err := day4.ReadFromFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

*/

/*
	day4.CopyFile("my_test_file.txt", "my_new_text_file.txt")
	day4.ReadFromFile("my_new_text_file.txt")
*/

/*
dir, _ := os.Getwd()
	//fmt.Println(dir)
	files, err := day4.ListFiles(dir)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
*/

/*
u := day4.User{Name: "Alice", Email: "alice@example.com", Age: 30, IsActive: true}
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	jsonBytes := string(data)
	fmt.Println(jsonBytes)

	jsonString := `{"name":"Bob","email":"bob@example.com","age":25,"is_active":false}`

	var newUser day4.User

	json.Unmarshal([]byte(jsonString), &newUser)

	fmt.Println(newUser)
*/

/*
users := []day4.User{
		{Name: "Alice", Email: "alice@example.com", Age: 30, IsActive: true},
		{Name: "Bob", Email: "bob@example.com", Age: 25, IsActive: false},
		{Name: "John", Email: "john@example.com", Age: 31, IsActive: true},
	}

	jsonBytes, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("users.json", jsonBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	var loadUsers []day4.User

	data, err := os.ReadFile("users.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &loadUsers)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(loadUsers)
*/
