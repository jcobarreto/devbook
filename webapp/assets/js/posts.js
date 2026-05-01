$('#new-post').on('submit', createPost);

$(document).on('click', '.like-post', likePost);
$(document).on('click', '.unlike-post', unlikePost);

$('#update-post').on('click', updatePost);

function createPost(event) {
  event.preventDefault();

  $.ajax({
    url: "/posts",
    method: "POST",
    data: {
      title: $('#title').val(),
      content: $('#content').val(),
    }
  }).done(function() {
    window.location = '/home';
  }).fail(function() {
    alert("Error creating post!");
  });
}

function likePost(event) {
  event.preventDefault();

  const clickedElement = $(event.currentTarget);
  const postID = clickedElement.closest('.p-5').data('post-id');

  clickedElement.prop('disabled', true);
  $.ajax({
    url: `/posts/${postID}/like`,
    method: "POST",
  }).done(function() {
    const likeCountElement = clickedElement.next('span');
    const likeCount = parseInt(likeCountElement.text());

    likeCountElement.text(likeCount + 1);

    clickedElement.addClass('unlike-post')
    clickedElement.addClass('text-danger')
    clickedElement.removeClass('like-post');

  }).fail(function() {
    alert("Error liking post!");
  }).always(function() {
    clickedElement.prop('disabled', false);
  });
}

function unlikePost(event) {
  event.preventDefault();

  const clickedElement = $(event.currentTarget);
  const postID = clickedElement.closest('.p-5').data('post-id');

  clickedElement.prop('disabled', true);
  $.ajax({
    url: `/posts/${postID}/unlike`,
    method: "POST",
  }).done(function() {
    const likeCountElement = clickedElement.next('span');
    const likeCount = parseInt(likeCountElement.text());

    likeCountElement.text(likeCount - 1);

    clickedElement.removeClass('unlike-post');
    clickedElement.removeClass('text-danger')
    clickedElement.addClass('like-post')

  }).fail(function() {
    alert("Error unliking post!");
  }).always(function() {
    clickedElement.prop('disabled', false);
  });
}

function updatePost(event) {
  $(this).prop('disabled', true);

  const postID = clickedElement.data('post-id');

  $.ajax({
    url: `/posts/${postID}`,
    method: "PUT",
    data: {
      title: $('#title').val(),
      content: $('#content').val(),
    }
  }).done(function() {
    alert("Post updated successfully!");
  }).fail(function() {
    alert("Error updating post!");
  }).always(function() {
    $('#update-post').prop('disabled', false);
  });
}
