package main

import "fmt"

const numDisks = 3

var totalMoves int

// Add a disk to the beginning of the post.
func push(post []int, disk int) []int {
	post = append(post, disk)
	return post
}

// Remove the first disk from the post.
// Return that disk and the revised post.
func pop(post []int) (int, []int) {
	item := post[len(post)-1]
	post = post[:len(post)-1]
	return item, post
}

// Move one disk from fromPost to toPost.
func moveDisk(posts [][]int, fromPost, toPost int) {
	var disk int
	disk, posts[fromPost] = pop(posts[fromPost])
	posts[toPost] = push(posts[toPost], disk)
}

// Draw the posts by showing the size of the disk at each level.
func drawPosts(posts [][]int) {
	for i, post := range posts {
		fmt.Print("post: ", i)
		fmt.Println(post)
	}
	fmt.Println("=======")
}

// Move the disks from fromPost to toPost
// using tempPost as temporary storage.
func moveDisks(posts [][]int, numToMove, fromPost, toPost, tempPost int) {
	// if we need to move more than one disks, we recursively call
	// using the tempPost now as the target
	if numToMove > 1 {
		moveDisks(posts, numToMove-1, fromPost, tempPost, toPost)
	}
	// here we only have to move one disk
	moveDisk(posts, fromPost, toPost)
	totalMoves += 1
	drawPosts(posts)
	// if we had to move more than one initially, then there are
	// some left in the tempPost
	if numToMove > 1 {
		moveDisks(posts, numToMove-1, tempPost, toPost, fromPost)
	}
}

func main() {
	// Make three posts.
	posts := [][]int{}

	// Push the disks onto post 0 biggest first.
	posts = append(posts, []int{})
	for disk := numDisks; disk > 0; disk-- {
		posts[0] = push(posts[0], disk)
	}

	// Make the other posts empty.
	for p := 1; p < 3; p++ {
		posts = append(posts, []int{})
	}

	// Draw the initial setup.
	drawPosts(posts)

	// Move the disks.
	moveDisks(posts, numDisks, 0, 1, 2)
	fmt.Println("Total moves: ", totalMoves)
}
