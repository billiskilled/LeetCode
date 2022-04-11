class Solution202204100201   {
    public static void main(String[] args) {
        System.out.println(deleteText("Singing dancing in the rain", 1));
    }
    public static String deleteText(String article, int index) {
        if (article == null || article.length() == 0 || index < 0 || index >= article.length()) {
            return article;
        }

        if (article.charAt(index) == ' ') {
            return article;
        }

        StringBuilder sb = new StringBuilder();
        StringBuilder cur = new StringBuilder();
        boolean find = false;
        for (int i = 0; i < article.length(); i++) {
            if (find) {
                sb.append(article.charAt(i));
                continue;
            }
            if (article.charAt(i) == ' ') {
                sb.append(cur);
                cur = new StringBuilder();
            }

            if (i == index) {
                while (i < article.length() && article.charAt(i) != ' ') {
                    i++;
                }

                cur = new StringBuilder();
                if (i == article.length()) {
                    break;
                } else {
                    sb.append(' ');
                }

                find = true;
            }

            cur.append(article.charAt(i));
        }

        return sb.toString().trim();
    }
}