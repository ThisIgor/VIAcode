package article

import (
	"knowledge/database"
)

const (
	TypeUnknown = iota
	TypeObject
	TypeEvent
	TypeOrganization
	TypeThink
	TypePerson
	TypeTherm
)

func InitArticleTypeGroup() {
	CreateArticleTypeGroup(TypeObject, "Объект", "Biological", "Объект живой природы (человек, фауна и флора) Биомеханический бобр")
	CreateArticleTypeGroup(TypeObject, "Объект", "Inanimate nature", "Объект неживой природы (минерал, планета, атмосферное явление)")
	CreateArticleTypeGroup(TypeObject, "Объект", "Geographical", "Природный географический объект (река, гора)")
	CreateArticleTypeGroup(TypeObject, "Объект", "Location", "Созданный, построенный географически локализованный объект: государство, город, крепость")
	CreateArticleTypeGroup(TypeObject, "Объект", "Culture", "Объект художественной культуры, произведение искусства, артефакт")
	CreateArticleTypeGroup(TypeObject, "Объект", "Fabricatio", "Объект технической культуры, изделие, устройство")
	CreateArticleTypeGroup(TypeEvent, "Событие", "Diana", "Природные явления, в том числе катастрофы")
	CreateArticleTypeGroup(TypeEvent, "Событие", "Ancient", "Археологическая эпоха, археологическая культура, древняя цивилизация")
	CreateArticleTypeGroup(TypeEvent, "Событие", "Quo-ad", "Исторические, экономические, цивилизационные процессы")
	CreateArticleTypeGroup(TypeEvent, "Событие", "Event", "Историческое событие, военная операция, спортивное состязание, фестиваль, конкурс")
	CreateArticleTypeGroup(TypeEvent, "Событие", "Event", "Технологии, технологические процессы")
	CreateArticleTypeGroup(TypeOrganization, "Организация", "Nation", "Этнос, нация, народ, этническая, этноконфессиональная группа, социальная группа")
	CreateArticleTypeGroup(TypeOrganization, "Организация", "Media", "СМИ, общественные и научные коммуникации, конференции")
	CreateArticleTypeGroup(TypeOrganization, "Организация", "Organization", "Предприятие, учреждение, институт, общественная организация, ведомство, заведение, трест, станция, база, фирма, концерн")
	CreateArticleTypeGroup(TypeThink, "Мысль", "Religion", "Философские, религиозные, богословские, социологические, психологические и др. учения, концепции")
	CreateArticleTypeGroup(TypeThink, "Мысль", "W", "Научное познание (математика, астрономия, лингвистика)")
	CreateArticleTypeGroup(TypeThink, "Мысль", "Jus", "Закон, договор, правило, конвенция, пакт")
	CreateArticleTypeGroup(TypeThink, "Мысль", "Humane", "Обычаи, игры, религиозные культы, праздники")
	CreateArticleTypeGroup(TypeThink, "Мысль", "Vis", "способы действия, влияния, регулирования (финансовые инструменты, политические и экономические методы, маркетинг, реклама, СМИ)")
	CreateArticleTypeGroup(TypeThink, "Мысль", "System", "Коммуникативная система (языки народов мира, языки программирования и др.)")
	CreateArticleTypeGroup(TypePerson, "Персона", "Person", "Персона")
	CreateArticleTypeGroup(TypeTherm, "Термин", "Term", "Термин")
}

func InitArticleType() {
	CreateArticleType("Ancient", "Археологическая эпоха, археологическая культура, древняя цивилизация", TypeUnknown, TypeUnknown)
	CreateArticleType("Biological", "Объект живой природы (человек, флора и фауна)", TypeUnknown, TypeUnknown)
	CreateArticleType("Culture", "Объект художественной культуры, произведение искусства, артефакт", TypeUnknown, TypeUnknown)
	CreateArticleType("Diana", "Природные явления, в том числе катастрофы", TypeUnknown, TypeUnknown)
	CreateArticleType("Event", "Историческое событие, военная операция, спортивное состязание, фестиваль, конкурс", TypeUnknown, TypeUnknown)
	CreateArticleType("Fabricatio", "Объект технической культуры, изделие, устройство", TypeUnknown, TypeUnknown)
	CreateArticleType("Geographical", "Природный географический объект (река, гора)", TypeUnknown, TypeUnknown)
	CreateArticleType("Humane", "Обычаи, игры, религиозные праздники", TypeUnknown, TypeUnknown)
	CreateArticleType("Inanimate nature", "Объект неживой природы (минерал, планета, атмосферное явление)", TypeUnknown, TypeUnknown)
	CreateArticleType("Jus", "Закон, договор, правило, конвенция, пакт", TypeUnknown, TypeUnknown)
	CreateArticleType("Location", "Созданный, построенный географически локализованный объект: государство, город, крепость", TypeUnknown, TypeUnknown)
	CreateArticleType("Media", "СМИ, общественные и научные коммуникации, конференции", TypeUnknown, TypeUnknown)
	CreateArticleType("Nation", "Этнос, нация, народ, этническая / этноконфессиональная группа, социальная группа", TypeUnknown, TypeUnknown)
	CreateArticleType("Organization", "Предприятие, учреждение, институт, общественная организация, ведомство, заведение, трест, станция, база, фирма, концерн, СМИ", TypeUnknown, TypeUnknown)
	CreateArticleType("Person", "Персона", TypeUnknown, TypeUnknown)
	CreateArticleType("Quo-ad", "Исторические, экономические, цивилизационные процессы", TypeUnknown, TypeUnknown)
	CreateArticleType("Religion", "Философские, религиозные, богословские, социологические, психологические и др. учения, концепции", TypeUnknown, TypeUnknown)
	CreateArticleType("System", "Коммуникативная система (языки народов мира, языки программирования и др.)", TypeUnknown, TypeUnknown)
	CreateArticleType("Term", "Термин", TypeUnknown, TypeUnknown)
	CreateArticleType("Utilis", "Технологии, технологические процессы", TypeUnknown, TypeUnknown)
	CreateArticleType("Vis", "Способы действия, влияния, регулирования (финансовые инструменты, политические и экономические методы, маркетинг, реклама, СМИ)", TypeUnknown, TypeUnknown)
	CreateArticleType("W", "Научное познание (математика, астрономия, лингвистика)", TypeUnknown, TypeUnknown)
	CreateArticleType("Z", "Прочее", TypeUnknown, TypeUnknown)
}

func InitArticles() {
	ClearArticles()
	InitArticleTypeGroup()
	InitArticleType()
}

func ClearArticles() {
	db := database.Connect()
	defer db.Close()
	db.Delete(ArticleType{})
	db.Delete(ArticleTypeGroup{})
}
